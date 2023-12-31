package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
	"time"
)

// ZapLogQueryHook is a pg.QueryHook that logs queries. Writes log
// to zap.Logger with a zap.DebugLevel
type ZapLogQueryHook struct {
	logger *zap.Logger
}

func NewZapLogQueryHook(logger *zap.Logger) pg.QueryHook {
	return &ZapLogQueryHook{logger}
}

func (h *ZapLogQueryHook) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (h *ZapLogQueryHook) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	var (
		dur = time.Now().Sub(q.StartTime)
	)

	query, err := q.FormattedQuery()

	if err == nil {
		h.logger.Debug(string(query), zap.Duration("dur", dur))
	}

	return nil
}
