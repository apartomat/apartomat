package main

import (
	"context"
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

type loggerHook struct {
	instance string
	logger   *zap.Logger
}

func (h loggerHook) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	query, err := q.FormattedQuery()

	if err == nil {
		h.logger.Debug(h.instance, zap.ByteString("query", query))
	}

	return c, nil
}

func (h loggerHook) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	return nil
}
