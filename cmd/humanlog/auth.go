package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/fatih/color"
	"github.com/humanlogio/api/go/svc/auth/v1/authv1connect"
	userpb "github.com/humanlogio/api/go/svc/user/v1"
	"github.com/humanlogio/api/go/svc/user/v1/userv1connect"
	"github.com/humanlogio/humanlog/internal/pkg/config"
	"github.com/humanlogio/humanlog/internal/pkg/state"
	"github.com/humanlogio/humanlog/pkg/auth"
	"github.com/pkg/browser"
	"github.com/urfave/cli"
)

const (
	authCmdName = "auth"
)

func authCmd(
	getCtx func(cctx *cli.Context) context.Context,
	getLogger func(cctx *cli.Context) *slog.Logger,
	getCfg func(cctx *cli.Context) *config.Config,
	getState func(cctx *cli.Context) *state.State,
	getTokenSource func(cctx *cli.Context) *auth.UserRefreshableTokenSource,
	getAPIUrl func(cctx *cli.Context) string,
	getHTTPClient func(cctx *cli.Context) *http.Client,
) cli.Command {
	return cli.Command{
		Name:  authCmdName,
		Usage: "Authenticate with humanlog.io",
		Subcommands: cli.Commands{
			{
				Name: "login",
				Action: func(cctx *cli.Context) error {
					ctx := getCtx(cctx)
					state := getState(cctx)
					tokenSource := getTokenSource(cctx)
					apiURL := getAPIUrl(cctx)
					httpClient := getHTTPClient(cctx)
					authClient := authv1connect.NewAuthServiceClient(httpClient, apiURL)
					_, err := performLoginFlow(ctx, state, authClient, tokenSource)
					return err
				},
			},
			{
				Name: "whoami",
				Action: func(cctx *cli.Context) error {
					ctx := getCtx(cctx)
					apiURL := getAPIUrl(cctx)
					state := getState(cctx)
					httpClient := getHTTPClient(cctx)
					tokenSource := getTokenSource(cctx)
					userToken, err := ensureLoggedIn(ctx, cctx, state, tokenSource, apiURL, httpClient)
					if err != nil {
						return fmt.Errorf("looking up local user state: %v", err)
					}

					ll := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
					clOpts := connect.WithInterceptors(
						auth.Interceptors(ll, tokenSource)...,
					)

					printFact("token user ID", userToken.UserId)

					userClient := userv1connect.NewUserServiceClient(httpClient, apiURL, clOpts)
					res, err := userClient.Whoami(ctx, connect.NewRequest(&userpb.WhoamiRequest{}))
					if err != nil {
						return fmt.Errorf("looking up who you are: %v", err)
					}

					printFact("email", res.Msg.User.Email)
					printFact("verified", res.Msg.User.EmailVerified)
					printFact("first name", res.Msg.User.FirstName)
					printFact("last name", res.Msg.User.LastName)
					printFact("registered since", res.Msg.User.CreatedAt.AsTime())
					printFact("logged into org", res.Msg.CurrentOrganization.Name)

					return nil
				},
			},
			{
				Name: "logout",
				Action: func(cctx *cli.Context) error {
					ctx := getCtx(cctx)
					apiURL := getAPIUrl(cctx)
					httpClient := getHTTPClient(cctx)
					tokenSource := getTokenSource(cctx)

					ll := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
					clOpts := connect.WithInterceptors(
						auth.Interceptors(ll, tokenSource)...,
					)
					userClient := userv1connect.NewUserServiceClient(httpClient, apiURL, clOpts)
					res, err := userClient.GetLogoutURL(ctx, connect.NewRequest(&userpb.GetLogoutURLRequest{}))
					if err != nil {
						return fmt.Errorf("retrieving logout URL")
					}
					if err := browser.OpenURL(res.Msg.GetLogoutUrl()); err != nil {
						return fmt.Errorf("opening logout URL")
					}
					return tokenSource.ClearToken(ctx)
				},
			},
		},
	}
}

func promptToLogin() {
	log.Print(
		color.YellowString("You are not logged in."),
	)
	log.Print(
		color.YellowString("Run `%s` to log in.", color.New(color.Bold).Sprint("humanlog auth login")),
	)
}
