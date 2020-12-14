package graphql

import (
	"context"
	"fmt"
	schemaorg "github.com/ztsu/apartomat/internal/pkg/schema.org"
)

type shoppingListResolver struct {
	*rootResolver
}

func (r *shoppingListResolver) FindProductOnPage(ctx context.Context, obj *ShoppingList, url string) (*Product, error) {
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
