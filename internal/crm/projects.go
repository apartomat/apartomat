package crm

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apartomat/apartomat/internal/crm/auth"
	. "github.com/apartomat/apartomat/internal/store/projects"
	sites "github.com/apartomat/apartomat/internal/store/public_sites"
	"github.com/apartomat/apartomat/internal/store/workspaces"
)

func (u *CRM) GetProject(ctx context.Context, id string) (*Project, error) {
	project, err := u.Projects.Get(ctx, IDIn(id))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanGetProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s): %w", project.ID, ErrForbidden)
	}

	return project, nil
}

func (u *CRM) CreateProject(
	ctx context.Context,
	workspaceID string,
	name string,
	startAt, endAt *time.Time,
) (*Project, error) {
	workspace, err := u.Workspaces.Get(ctx, workspaces.IDIn(workspaceID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanCreateProject(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't create project in workspace (id=%s): %w", workspace.ID, ErrForbidden)
	}

	project := NewProject(MustGenerateNanoID(), name, startAt, endAt, workspaceID)

	if err := u.Projects.Save(ctx, project); err != nil {
		return nil, err
	}

	siteID := MustGenerateNanoID()

	site := sites.NewPublicSite(
		siteID,
		fmt.Sprintf("https://p.apartomat.ru/%s", siteID),
		sites.StatusNotPublic,
		sites.PublicSiteSettings{
			AllowVisualizations: false,
			AllowAlbums:         false,
		},
		project.ID,
	)

	if err := u.PublicSites.Save(ctx, site); err != nil {
		return nil, err
	}

	return project, nil
}

func (u *CRM) ChangeProjectStatus(ctx context.Context, projectID string, status Status) (*Project, error) {
	project, err := u.Projects.Get(ctx, IDIn(projectID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update project (id=%s): %w", project.ID, ErrForbidden)
	}

	project.ChangeStatus(status)

	if err := u.Projects.Save(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (u *CRM) ChangeProjectDates(ctx context.Context, projectID string, startAt, endAt *time.Time) (*Project, error) {
	project, err := u.Projects.Get(ctx, IDIn(projectID))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't update project (id=%s): %w", project.ID, ErrForbidden)
	}

	project.ChangeDates(startAt, endAt)

	if err := u.Projects.Save(ctx, project); err != nil {
		return nil, err
	}

	return project, nil
}

func (u *CRM) GetProjectPublicSite(ctx context.Context, projectId string) (*sites.PublicSite, error) {
	if ok, err := u.Acl.CanGetPublicSiteOfProjectID(ctx, auth.UserFromCtx(ctx), projectId); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) public site: %w", projectId, ErrForbidden)
	}

	return u.PublicSites.Get(ctx, sites.ProjectIDIn(projectId))
}

func (u *CRM) MakeProjectPublic(ctx context.Context, projectId string) (*sites.PublicSite, error) {
	proj, err := u.Projects.Get(ctx, IDIn(projectId))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateProject(ctx, auth.UserFromCtx(ctx), proj); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't make project (id=%s) public: %w", proj.ID, ErrForbidden)
	}

	site, err := u.PublicSites.Get(ctx, sites.ProjectIDIn(proj.ID))
	if errors.Is(err, sites.ErrPublicSiteNotFound) {
		siteID := MustGenerateNanoID()

		s := sites.NewPublicSite(
			siteID,
			fmt.Sprintf("https://p.apartomat.ru/%s", siteID),
			sites.StatusNotPublic,
			sites.PublicSiteSettings{
				AllowVisualizations: true,
				AllowAlbums:         true,
			},
			proj.ID,
		)

		site = &s

	} else if err != nil {
		return nil, err
	}

	if err := site.ToPublic(); err != nil {
		return nil, fmt.Errorf("can't make project public: %w", err)
	}

	if err := u.PublicSites.Save(ctx, *site); err != nil {
		return nil, err
	}

	return site, nil
}

func (u *CRM) MakeProjectNotPublic(ctx context.Context, projectId string) (*sites.PublicSite, error) {
	proj, err := u.Projects.Get(ctx, IDIn(projectId))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateProject(ctx, auth.UserFromCtx(ctx), proj); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't make project (id=%s) not public: %w", proj.ID, ErrForbidden)
	}

	site, err := u.PublicSites.Get(ctx, sites.ProjectIDIn(proj.ID))
	if errors.Is(err, sites.ErrPublicSiteNotFound) {
		siteID := MustGenerateNanoID()

		s := sites.NewPublicSite(
			siteID,
			fmt.Sprintf("https://p.apartomat.ru/%s", siteID),
			sites.StatusPublic,
			sites.PublicSiteSettings{
				AllowVisualizations: true,
				AllowAlbums:         true,
			},
			proj.ID,
		)

		site = &s

	} else if err != nil {
		return nil, err
	}

	if err := site.ToNotPublic(); err != nil {
		return nil, fmt.Errorf("can't make project not public: %w", err)
	}

	if err := u.PublicSites.Save(ctx, *site); err != nil {
		return nil, err
	}

	return site, nil
}
