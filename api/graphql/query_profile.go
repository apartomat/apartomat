package graphql

import (
	"context"
	"errors"

	apartomat "github.com/apartomat/apartomat/internal"
	"github.com/apartomat/apartomat/internal/auth"
	"github.com/apartomat/apartomat/internal/pkg/gravatar"
	"go.uber.org/zap"
)

func (r *queryResolver) Profile(ctx context.Context) (UserProfileResult, error) {
	if userCtx := auth.UserFromCtx(ctx); userCtx != nil {
		user, err := r.useCases.GetUserProfile(ctx, userCtx.ID)
		if err != nil {
			if errors.Is(err, apartomat.ErrForbidden) {
				return forbidden()
			}

			r.logger.Error("can't get profile", zap.String("user", userCtx.ID), zap.Error(err))

			return serverError()
		}

		var (
			grava *Gravatar
		)

		if user.UseGravatar {
			grava = &Gravatar{URL: gravatar.Url(user.Email)}
		}

		profile := UserProfile{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
			Gravatar: grava,
			Abbr:     userAbbr(user.FullName, user.Email),
		}

		if user.DefaultWorkspaceID != nil {
			profile.DefaultWorkspace = &Workspace{
				ID: *user.DefaultWorkspaceID,
			}
		}

		return profile, nil
	}

	return forbidden()
}
