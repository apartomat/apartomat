package dataloaders

import "context"

const (
	DataLoadersContextKey = "dataloaders"
)

func WithDataLoaders(ctx context.Context, dataLoaders *DataLoaders) context.Context {
	return context.WithValue(ctx, DataLoadersContextKey, dataLoaders)
}

func FromContext(ctx context.Context) *DataLoaders {
	return ctx.Value(DataLoadersContextKey).(*DataLoaders)
}
