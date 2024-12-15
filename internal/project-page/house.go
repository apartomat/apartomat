package project_page

import (
	"context"
	"github.com/apartomat/apartomat/internal/store/houses"
	"github.com/apartomat/apartomat/internal/store/projectpage"
)

func (u *Service) GetHouse(ctx context.Context, projectPageID string) (*houses.House, error) {
	page, err := u.ProjectPages.Get(ctx, projectpage.IDIn(projectPageID))
	if err != nil {
		return nil, err
	}

	house, err := u.Houses.Get(ctx, houses.ProjectIDIn(page.ProjectID))
	if err != nil {
		return nil, err
	}

	return house, nil
}
