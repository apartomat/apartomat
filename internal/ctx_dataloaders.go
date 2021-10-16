package apartomat

import (
	"context"
)

const dataLoadersCtxKey = "dataloaders"

type DataLoaders struct{}

func WithDataLoadersCtx(ctx context.Context, dataLoaders *DataLoaders) context.Context {
	return context.WithValue(ctx, dataLoadersCtxKey, dataLoaders)
}

func DataLoadersFromCtx(ctx context.Context) *DataLoaders {
	return ctx.Value(dataLoadersCtxKey).(*DataLoaders)
}
