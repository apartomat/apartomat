package graphql

import (
	"context"
	"fmt"
	schemaorg "github.com/apartomat/apartomat/internal/pkg/schema.org"
)

func (r *rootResolver) ShoppinglistQuery() ShoppinglistQueryResolver {
	return &shoppinglistQueryResolver{r}
}

type shoppinglistQueryResolver struct {
	*rootResolver
}

func (r *shoppinglistQueryResolver) ProductOnPage(ctx context.Context, obj *ShoppinglistQuery, url string) (*Product, error) {
	p, err := schemaorg.FindProductCtx(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("can't find product on page: %w", err)
	}

	return productToGraphQL(p), nil
}

func productToGraphQL(p *schemaorg.Product) *Product {
	if p == nil {
		return nil
	}

	var (
		img string
	)
	if len(p.Image) > 0 {
		img = p.Image[0]
	}

	return &Product{Name: p.Name, Description: p.Description, Image: img}
}

//

func (r *queryResolver) Shoppinglist(ctx context.Context) (*ShoppinglistQuery, error) {
	return &ShoppinglistQuery{}, nil
}
