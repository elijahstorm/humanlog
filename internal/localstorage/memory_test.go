package localstorage

import (
	"context"
	"testing"
	"time"

	typesv1 "github.com/humanlogio/api/go/types/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMemoryStorage(t *testing.T) {
	tests := []struct {
		name  string
		q     *typesv1.LogQuery
		input []*typesv1.LogEventGroup
		want  []*typesv1.LogEventGroup
	}{
		{
			name: "nothing",
			q: &typesv1.LogQuery{
				To: timestamppb.New(musttime("2006-01-02T15:04:06.000")),
			},
			input: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.001")), Raw: []byte("hello world 1")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.004")), Raw: []byte("hello world 4")},
					},
				},
			},
			want: []*typesv1.LogEventGroup{},
		},
		{
			name: "all",
			q: &typesv1.LogQuery{
				To: timestamppb.New(musttime("2006-01-02T15:04:06.005")),
			},
			input: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.001")), Raw: []byte("hello world 1")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.004")), Raw: []byte("hello world 4")},
					},
				},
			},
			want: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.001")), Raw: []byte("hello world 1")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.004")), Raw: []byte("hello world 4")},
					},
				},
			},
		},
		{
			name: "skip last",
			q: &typesv1.LogQuery{
				To: timestamppb.New(musttime("2006-01-02T15:04:06.004")),
			},
			input: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.001")), Raw: []byte("hello world 1")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.004")), Raw: []byte("hello world 4")},
					},
				},
			},
			want: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.001")), Raw: []byte("hello world 1")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
					},
				},
			},
		},
		{
			name: "skip first",
			q: &typesv1.LogQuery{
				From: timestamppb.New(musttime("2006-01-02T15:04:06.002")),
				To:   timestamppb.New(musttime("2006-01-02T15:04:06.005")),
			},
			input: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.001")), Raw: []byte("hello world 1")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.004")), Raw: []byte("hello world 4")},
					},
				},
			},
			want: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.004")), Raw: []byte("hello world 4")},
					},
				},
			},
		},
		{
			name: "slice",
			q: &typesv1.LogQuery{
				From: timestamppb.New(musttime("2006-01-02T15:04:06.002")),
				To:   timestamppb.New(musttime("2006-01-02T15:04:06.004")),
			},
			input: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.001")), Raw: []byte("hello world 1")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.004")), Raw: []byte("hello world 4")},
					},
				},
			},
			want: []*typesv1.LogEventGroup{
				{
					MachineId: 1, SessionId: 2,
					Logs: []*typesv1.LogEvent{
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.002")), Raw: []byte("hello world 2")},
						{ParsedAt: timestamppb.New(musttime("2006-01-02T15:04:06.003")), Raw: []byte("hello world 3")},
					},
				},
			},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mem := NewMemStorage()

			for _, leg := range tt.input {
				snk, _, err := mem.SinkFor(leg.MachineId, leg.SessionId)
				require.NoError(t, err)
				for _, ev := range leg.Logs {
					err = snk.Receive(ctx, ev)
					require.NoError(t, err)
				}
				err = snk.Flush(ctx)
				require.NoError(t, err)
			}

			cursors, err := mem.Query(ctx, tt.q)
			require.NoError(t, err)
			got := drainCursors(t, ctx, cursors)

			require.Len(t, got, len(tt.want))
			for i := range tt.want {
				require.Equal(t, protojson.Format(tt.want[i]), protojson.Format(got[i]))
			}
		})
	}
}

func drainCursors(t *testing.T, ctx context.Context, cursors []Cursor) []*typesv1.LogEventGroup {
	out := make([]*typesv1.LogEventGroup, 0, len(cursors))
	for _, cursor := range cursors {
		mid, sid := cursor.IDs()
		leg := &typesv1.LogEventGroup{
			MachineId: mid, SessionId: sid,
		}
		for cursor.Next(ctx) {
			leg.Logs = append(leg.Logs, cursor.Event())
		}
		require.NoError(t, cursor.Err())
		out = append(out, leg)
	}
	return out
}

func musttime(str string) time.Time {
	t, err := time.Parse("2006-01-02T15:04:05.000", str)
	if err != nil {
		panic(err)
	}
	return t
}
