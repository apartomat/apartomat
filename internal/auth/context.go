package auth

import "context"

const userCtxKey = "user"

type UserCtx struct {
	ID    string
	Email string
}

func WithUserCtx(ctx context.Context, userCtx *UserCtx) context.Context {
	return context.WithValue(ctx, userCtxKey, userCtx)
}

func UserFromCtx(ctx context.Context) *UserCtx {
	user, _ := ctx.Value(userCtxKey).(*UserCtx)
	return user
}
