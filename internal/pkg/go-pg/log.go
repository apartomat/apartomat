package go_pg

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log/slog"
	"time"
)

type LogQueryHook struct {
	logger *slog.Logger
}

func NewLogQueryHook(logger *slog.Logger) pg.QueryHook {
	return &LogQueryHook{logger}
}

func (h *LogQueryHook) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (h *LogQueryHook) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	var (
		dur = time.Now().Sub(q.StartTime)
	)

	query, err := q.FormattedQuery()

	if err == nil {
		h.logger.DebugContext(
			ctx,
			"Database query",
			slog.Duration("dur", dur),
			slog.String("query", string(query)),
			slog.String("ctx", QueryContext(ctx)),
		)
	}

	return nil
}
