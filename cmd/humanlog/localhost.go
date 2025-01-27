package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"syscall"
	"time"

	connectcors "connectrpc.com/cors"

	"github.com/humanlogio/api/go/svc/ingest/v1/ingestv1connect"
	"github.com/humanlogio/api/go/svc/localhost/v1/localhostv1connect"
	"github.com/humanlogio/api/go/svc/query/v1/queryv1connect"
	typesv1 "github.com/humanlogio/api/go/types/v1"
	"github.com/humanlogio/humanlog/internal/localstorage"
	"github.com/humanlogio/humanlog/internal/localsvc"
	"github.com/humanlogio/humanlog/internal/pkg/config"
	"github.com/humanlogio/humanlog/internal/pkg/state"
	"github.com/humanlogio/humanlog/pkg/sink"
	"github.com/humanlogio/humanlog/pkg/sink/logsvcsink"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sys/unix"
)

func isEADDRINUSE(err error) bool {
	nerr, ok := err.(*net.OpError)
	if !ok {
		return false
	}
	nserr, ok := nerr.Err.(*os.SyscallError)
	if !ok {
		return false
	}
	if nserr.Syscall != "bind" {
		return false
	}
	nserrno, ok := nserr.Err.(syscall.Errno)
	if !ok {
		return false
	}
	return nserrno == unix.EADDRINUSE
}

func startLocalhostServer(
	ctx context.Context,
	ll *slog.Logger,
	cfg *config.Config,
	state *state.State,
	machineID uint64,
	port int,
	localhostHttpClient *http.Client,
	ownVersion *typesv1.Version,
) (localsink sink.Sink, done func(context.Context) error, err error) {
	localhostAddr := net.JoinHostPort("localhost", strconv.Itoa(port))

	l, err := net.Listen("tcp", localhostAddr)
	if err != nil && !isEADDRINUSE(err) {
		return nil, nil, fmt.Errorf("listening on host/port: %v", err)
	}
	if isEADDRINUSE(err) {
		// TODO(antoine):
		// 1) log to localhost until it's gone
		// 2) try to gain the socket, if fail; goto 1)
		// 3) serve the localhost service + save local logs + forward them to remote
		addr, err := url.Parse("http://" + localhostAddr)
		if err != nil {
			panic(err)
		}
		log.Printf("DEBUG: sending logs to localhost forwarder")
		client := ingestv1connect.NewIngestServiceClient(localhostHttpClient, addr.String())
		localhostSink := logsvcsink.StartStreamSink(ctx, client, "local", machineID, 1<<20, 100*time.Millisecond, true)
		return localhostSink, func(ctx context.Context) error {
			return localhostSink.Flush(ctx)
		}, nil
	}
	storage := localstorage.NewMemStorage()
	ownSink, _, err := storage.SinkFor(int64(machineID), time.Now().UnixNano())
	if err != nil {
		return nil, nil, fmt.Errorf("can't create own sink: %v", err)
	}

	mux := http.NewServeMux()

	localhostsvc := localsvc.New(ll, state, ownVersion, storage)
	mux.Handle(localhostv1connect.NewLocalhostServiceHandler(localhostsvc))
	mux.Handle(ingestv1connect.NewIngestServiceHandler(localhostsvc))
	mux.Handle(queryv1connect.NewQueryServiceHandler(localhostsvc))

	hdl := h2c.NewHandler(mux, &http2.Server{})
	hdl = withCORS(hdl)

	srv := http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hdl.ServeHTTP(w, r)
	})}

	go func() {
		log.Printf("localhost service available on %s, visit `https://humanlog.io` so see your logs", l.Addr().String())
		if err := srv.Serve(l); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("failed to serve localhost service, giving up: %v", err)
		}
	}()

	return ownSink, func(ctx context.Context) error {
		errc := make(chan error, 2)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			errc <- srv.Shutdown(ctx)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			errc <- ownSink.Flush(ctx)
		}()
		wg.Wait()
		close(errc)
		l.Close()
		var ferr error
		for err := range errc {
			if ferr == nil {
				ferr = err
			} else {
				ll.ErrorContext(ctx, "multiple errors", slog.Any("err", err))
			}
		}
		return ferr
	}, nil
}

// withCORS adds CORS support to a Connect HTTP handler.
func withCORS(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		// Debug: true,
		AllowedOrigins: []string{
			"https://humanlog.io",
			"https://humanlog.dev",
			"https://app.humanlog.dev",
			"https://app.humanlog.dev:3000",
			"https://humanlog.sh",
			"http://localhost:3000",
			"https://humanlog.test:3000",
		},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}
