package postgres

import (
	"context"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

// loggerHook is a pg.QueryHook that logs queries. Writes log
// to zap.Logger with a zap.DebugLevel
type loggerHook struct {
	logger *zap.Logger
}

func NewLoggerHook(logger *zap.Logger) pg.QueryHook {
	return &loggerHook{logger}
}

func (h *loggerHook) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	query, err := q.FormattedQuery()

	if err == nil {
		h.logger.Debug(string(query))
	}

	return c, nil
}

func (h *loggerHook) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}
