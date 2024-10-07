package log

import (
	"context"
	"log/slog"
)

//

type attrsLogContext string

var (
	attrsLogContextKey attrsLogContext = "slog_attrs"
)

func WithLogAttr(parent context.Context, attrs ...slog.Attr) context.Context {
	if v, ok := parent.Value(attrsLogContextKey).([]slog.Attr); ok {
		return context.WithValue(parent, attrsLogContextKey, append(v, attrs...))
	}

	return context.WithValue(parent, attrsLogContextKey, attrs)
}

//

var _ slog.Handler = (*AttrHandler)(nil)

type AttrHandler struct {
	next slog.Handler
}

func NewAttrHandler(next slog.Handler) *AttrHandler {
	return &AttrHandler{next}
}

func (h *AttrHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

func (h *AttrHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &AttrHandler{h.next.WithAttrs(attrs)}
}

func (h *AttrHandler) WithGroup(name string) slog.Handler {
	return &AttrHandler{h.next.WithGroup(name)}
}

func (h *AttrHandler) Handle(ctx context.Context, r slog.Record) error {
	if v, ok := ctx.Value(attrsLogContextKey).([]slog.Attr); ok {
		r.AddAttrs(v...)
	}

	return h.next.Handle(ctx, r)
}
