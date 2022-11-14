package dataloader

import (
	"context"
	"errors"
)

const dataLoadersCtxKey = "dataloaders"

type DataLoaders struct {
	Users *UserLoader
}

func WithDataLoadersCtx(ctx context.Context, dataLoaders *DataLoaders) context.Context {
	return context.WithValue(ctx, dataLoadersCtxKey, dataLoaders)
}

func DataLoadersFromCtx(ctx context.Context) *DataLoaders {
	return ctx.Value(dataLoadersCtxKey).(*DataLoaders)
}

func UserLoaderFromCtx(ctx context.Context) (*UserLoader, error) {
	loaders := DataLoadersFromCtx(ctx)
	if loaders == nil {
		return nil, errors.New("there are no data loaders in context")
	}

	if loaders.Users == nil {
		return nil, errors.New("there are no users data loader")
	}

	return loaders.Users, nil
}
