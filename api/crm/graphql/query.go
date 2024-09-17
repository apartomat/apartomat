package graphql

import "context"

var (
	Version = ""
)

type queryResolver struct {
	*rootResolver
}

func (r *queryResolver) Version(ctx context.Context) (string, error) {
	return Version, nil
}

func (r *queryResolver) Db(ctx context.Context) (bool, error) {
	var (
		err = r.db.PingContext(ctx)
	)

	return err == nil, nil
}

func (r *rootResolver) Query() QueryResolver { return &queryResolver{r} }
