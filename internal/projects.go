package apartomat

import (
	"context"
	"errors"
	"fmt"
	"github.com/apartomat/apartomat/internal/auth"
	. "github.com/apartomat/apartomat/internal/store/projects"
	sites "github.com/apartomat/apartomat/internal/store/public_sites"
	"github.com/apartomat/apartomat/internal/store/workspace_users"
	"github.com/apartomat/apartomat/internal/store/workspaces"
	"time"
)

func (u *Apartomat) CreateProject(
	ctx context.Context,
	workspaceID string,
	name string,
	startAt, endAt *time.Time,
) (*Project, error) {
	ws, err := u.Workspaces.List(ctx, workspaces.IDIn(workspaceID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(ws) == 0 {
		return nil, fmt.Errorf("workspace (id=%s): %w", workspaceID, ErrNotFound)
	}

	var (
		workspace = ws[0]
	)

	if ok, err := u.CanCreateProject(ctx, auth.UserFromCtx(ctx), workspace); err != nil {
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

func (u *Apartomat) CanCreateProject(ctx context.Context, subj *auth.UserCtx, obj *workspaces.Workspace) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(obj.ID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return true, nil
}

func (u *Apartomat) ChangeProjectStatus(ctx context.Context, projectID string, status Status) (*Project, error) {
	projects, err := u.Projects.List(ctx, IDIn(projectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanUpdateProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
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

func (u *Apartomat) ChangeProjectDates(ctx context.Context, projectID string, startAt, endAt *time.Time) (*Project, error) {
	projects, err := u.Projects.List(ctx, IDIn(projectID), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectID, ErrNotFound)
	}

	var (
		project = projects[0]
	)

	if ok, err := u.CanUpdateProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
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

func (u *Apartomat) CanUpdateProject(ctx context.Context, subj *auth.UserCtx, obj *Project) (bool, error) {
	if subj == nil {
		return false, nil
	}

	wu, err := u.WorkspaceUsers.List(
		ctx,
		workspace_users.And(
			workspace_users.WorkspaceIDIn(obj.WorkspaceID),
			workspace_users.UserIDIn(subj.ID),
		),
		1,
		0,
	)
	if err != nil {
		return false, err
	}

	if len(wu) == 0 {
		return false, nil
	}

	return wu[0].UserID == subj.ID, nil
}

func (u *Apartomat) GetProject(ctx context.Context, id string) (*Project, error) {
	prjs, err := u.Projects.List(ctx, IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", id, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanGetProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s): %w", project.ID, ErrForbidden)
	}

	return project, nil
}

func (u *Apartomat) CanGetProject(ctx context.Context, subj *auth.UserCtx, obj *Project) (bool, error) {
	return u.isProjectUser(ctx, subj, obj)
}

func (u *Apartomat) PublishProject(ctx context.Context, id string) (*Project, error) {
	prjs, err := u.Projects.List(ctx, IDIn(id), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", id, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanGetProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s): %w", project.ID, ErrForbidden)
	}

	return project, nil
}

func (u *Apartomat) GetProjectPublicSite(ctx context.Context, projectId string) (*sites.PublicSite, error) {
	prjs, err := u.Projects.List(ctx, IDIn(projectId), 1, 0)
	if err != nil {
		return nil, err
	}

	if len(prjs) == 0 {
		return nil, fmt.Errorf("project (id=%s): %w", projectId, ErrNotFound)
	}

	var (
		project = prjs[0]
	)

	if ok, err := u.CanGetProject(ctx, auth.UserFromCtx(ctx), project); err != nil {
		return nil, err
	} else if !ok {
		return nil, fmt.Errorf("can't get project (id=%s) public site: %w", project.ID, ErrForbidden)
	}

	return u.PublicSites.Get(ctx, sites.ProjectIDIn(project.ID))
}

func (u *Apartomat) MakeProjectPublic(ctx context.Context, projectId string) (*sites.PublicSite, error) {
	proj, err := u.Projects.Get(ctx, IDIn(projectId))
	if err != nil {
		return nil, err
	}

	if ok, err := u.CanUpdateProject(ctx, auth.UserFromCtx(ctx), proj); err != nil {
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

func (u *Apartomat) MakeProjectNotPublic(ctx context.Context, projectId string) (*sites.PublicSite, error) {
	proj, err := u.Projects.Get(ctx, IDIn(projectId))
	if err != nil {
		return nil, err
	}

	if ok, err := u.CanUpdateProject(ctx, auth.UserFromCtx(ctx), proj); err != nil {
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
