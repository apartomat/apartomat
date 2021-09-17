package graphql

import (
	"context"
	"github.com/apartomat/apartomat/internal/pkg/gravatar"
	"strings"
)

func (r *rootResolver) WorkspaceUser() WorkspaceUserResolver {
	return &workspaceUserResolver{r}
}

type workspaceUserResolver struct {
	*rootResolver
}

func (r *workspaceUserResolver) Profile(ctx context.Context, obj *WorkspaceUser) (*WorkspaceUserProfile, error) {
	user, err := r.useCases.GetWorkspaceUserProfile.Do(ctx, obj.Workspace.ID, obj.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	var (
		grava *Gravatar
	)

	if user.UseGravatar {
		grava = &Gravatar{URL: gravatar.Url(user.Email)}
	}

	return &WorkspaceUserProfile{
		ID:       obj.ID,
		Email:    user.Email,
		Gravatar: grava,
		FullName: user.FullName,
		Abbr:     abbr(user.FullName),
	}, nil
}

func abbr(str string) string {
	spl := strings.Split(str, " ")

	if len(spl) > 2 {
		spl = spl[0:2]
	}

	f := make([]string, len(spl))

	for i, s := range spl {
		sr := []rune(s)
		if len(sr) > 0 {
			f[i] = string(sr[0:1])
		} else {
			f[i] = s
		}
	}

	return strings.Join(f, "")
}
