package graphql

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"

	"github.com/apartomat/apartomat/internal/pkg/gravatar"
)

type workspaceUserResolver struct {
	*rootResolver
}

func (r *rootResolver) WorkspaceUser() WorkspaceUserResolver {
	return &workspaceUserResolver{r}
}

func (r *workspaceUserResolver) Profile(ctx context.Context, obj *WorkspaceUser) (*UserProfile, error) {
	user, err := r.crm.GetWorkspaceUserProfileDl(ctx, obj.Workspace.ID, obj.ID)
	if err != nil {
		slog.ErrorContext(ctx, "can't resolve workspace user profile", slog.String("user", obj.ID), slog.Any("err", err))

		return nil, errors.New("internal server error")
	}

	if user == nil {
		slog.ErrorContext(
			ctx,
			"can't resolve workspace user profile",
			slog.String("user", obj.ID),
			slog.Any("err", fmt.Errorf("user (id=%s) not found", obj.ID)),
		)

		return nil, errors.New("internal server error")
	}

	var (
		grava *Gravatar
	)

	if user.UseGravatar {
		grava = &Gravatar{URL: gravatar.Url(user.Email)}
	}

	profile := &UserProfile{
		ID:       obj.ID,
		Email:    user.Email,
		Gravatar: grava,
		FullName: user.FullName,
		Abbr:     userAbbr(user.FullName, user.Email),
	}

	if user.DefaultWorkspaceID != nil {
		profile.DefaultWorkspace = &Workspace{
			ID: *user.DefaultWorkspaceID,
		}
	}

	return profile, nil
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

func userAbbr(name, email string) string {
	if ab := abbr(name); ab != "" {
		return strings.ToUpper(ab)
	}

	return strings.ToUpper(email[0:2])
}
