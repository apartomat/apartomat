package graphql

import "context"

func (r *rootResolver) Workspace() WorkspaceResolver {
	return &workspaceResolver{r}
}

type workspaceResolver struct {
	*rootResolver
}

func (r *workspaceResolver) Projects(ctx context.Context, obj *Workspace) (*WorkspaceProjects, error) {
	return &WorkspaceProjects{}, nil
}

func (r *workspaceResolver) Users(ctx context.Context, obj *Workspace) (*WorkspaceUsers, error) {
	return &WorkspaceUsers{}, nil
}
