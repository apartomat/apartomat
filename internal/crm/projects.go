package crm

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/apartomat/apartomat/internal/store/projectpages"
	. "github.com/apartomat/apartomat/internal/store/projects"
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

	pageID := MustGenerateNanoID()

	page := projectpages.NewProjectPage(
		pageID,
		name,
		"",
		u.ProjectPageURL(pageID),
		projectpages.StatusNotPublic,
		projectpages.Settings{
			AllowVisualizations: true,
			AllowAlbums:         true,
		},
		project.ID,
	)

	if err := u.ProjectPages.Save(ctx, page); err != nil {
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

func (u *CRM) GetProjectPage(ctx context.Context, projectId string) (*projectpages.ProjectPage, error) {
	if ok, err := u.Acl.CanGetPublicSiteOfProjectID(ctx, auth.UserFromCtx(ctx), projectId); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) page: %w", projectId, ErrForbidden)
	}

	page, err := u.ProjectPages.Get(ctx, projectpages.ProjectIDIn(projectId))
	if err != nil {
		return nil, err
	}

	{
		baseURL := u.Params.ProjectPageBaseURL

		pageURL, err := url.Parse(page.URL)
		if err != nil {
			return nil, err
		}

		fixedPageURL := &url.URL{
			Scheme: baseURL.Scheme,
			Host:   baseURL.Host,
			Path:   pageURL.RequestURI(),
		}

		page.URL = fixedPageURL.String()
	}

	return page, nil
}

func (u *CRM) MakeProjectPublic(ctx context.Context, projectId string) (*projectpages.ProjectPage, error) {
	proj, err := u.Projects.Get(ctx, IDIn(projectId))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateProject(ctx, auth.UserFromCtx(ctx), proj); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't make project (id=%s) public: %w", proj.ID, ErrForbidden)
	}

	page, err := u.ProjectPages.Get(ctx, projectpages.ProjectIDIn(proj.ID))
	if errors.Is(err, projectpages.ErrProjectPageNotFound) {
		pageID := MustGenerateNanoID()

		p := projectpages.NewProjectPage(
			pageID,
			proj.Name,
			"",
			u.ProjectPageURL(pageID),
			projectpages.StatusNotPublic,
			projectpages.Settings{
				AllowVisualizations: true,
				AllowAlbums:         true,
			},
			proj.ID,
		)

		page = &p

	} else if err != nil {
		return nil, err
	}

	if err := page.ToPublic(); err != nil {
		return nil, fmt.Errorf("can't make project public: %w", err)
	}

	if err := u.ProjectPages.Save(ctx, *page); err != nil {
		return nil, err
	}

	return page, nil
}

func (u *CRM) MakeProjectNotPublic(ctx context.Context, projectId string) (*projectpages.ProjectPage, error) {
	proj, err := u.Projects.Get(ctx, IDIn(projectId))
	if err != nil {
		return nil, err
	}

	if ok, err := u.Acl.CanUpdateProject(ctx, auth.UserFromCtx(ctx), proj); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't make project (id=%s) not public: %w", proj.ID, ErrForbidden)
	}

	page, err := u.ProjectPages.Get(ctx, projectpages.ProjectIDIn(proj.ID))
	if errors.Is(err, projectpages.ErrProjectPageNotFound) {
		pageID := MustGenerateNanoID()

		p := projectpages.NewProjectPage(
			pageID,
			proj.Name,
			"",
			u.ProjectPageURL(pageID),
			projectpages.StatusPublic,
			projectpages.Settings{
				AllowVisualizations: true,
				AllowAlbums:         true,
			},
			proj.ID,
		)

		page = &p

	} else if err != nil {
		return nil, err
	}

	if err := page.ToNotPublic(); err != nil {
		return nil, fmt.Errorf("can't make project not public: %w", err)
	}

	if err := u.ProjectPages.Save(ctx, *page); err != nil {
		return nil, err
	}

	return page, nil
}

func (u *CRM) ProjectPageURL(pageID string) string {
	baseURL := u.Params.ProjectPageBaseURL

	pageURL := &url.URL{
		Scheme: baseURL.Scheme,
		Host:   baseURL.Host,
		Path:   pageID,
	}

	return pageURL.String()
}
