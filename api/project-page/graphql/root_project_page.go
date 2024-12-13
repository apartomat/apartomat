package graphql

import (
	"context"
	"errors"
	project_page "github.com/apartomat/apartomat/internal/project-page"
	"github.com/apartomat/apartomat/internal/store"
	"github.com/apartomat/apartomat/internal/store/houses"
)

func (r *rootResolver) ProjectPage() ProjectPageResolver {
	return &projectPageResolver{r}
}

type projectPageResolver struct {
	*rootResolver
}

func (r *projectPageResolver) House(ctx context.Context, obj *ProjectPage) (ProjectPageHouseResult, error) {
	house, err := r.projectPage.GetHouse(ctx, obj.ID)
	if err != nil {
		if errors.Is(err, project_page.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, store.ErrNotFound) {
			return notFound()
		}
	}

	return houseToGraphQL(house), nil
}

func houseToGraphQL(house *houses.House) *House {
	return &House{
		ID:             house.ID,
		City:           house.City,
		Address:        house.Address,
		HousingComplex: house.HousingComplex,
	}
}

func (r *projectPageResolver) Visualizations(ctx context.Context, obj *ProjectPage) (*ProjectPageVisualizations, error) {
	return &ProjectPageVisualizations{}, nil
}

func (r *projectPageResolver) Album(ctx context.Context, obj *ProjectPage) (ProjectPageAlbumResult, error) {
	album, file, err := r.projectPage.GetAlbumAndFile(ctx, obj.ID)
	if err != nil {
		if errors.Is(err, project_page.ErrForbidden) {
			return forbidden()
		}

		if errors.Is(err, store.ErrNotFound) {
			return notFound()
		}
	}

	return Album{
		ID:   album.ID,
		Name: album.Name,
		URL:  file.URL,
		Size: int(file.Size),
	}, nil
}
