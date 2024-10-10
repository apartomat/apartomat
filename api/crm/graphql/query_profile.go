package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
	"github.com/apartomat/apartomat/internal/crm/auth"
	"github.com/apartomat/apartomat/internal/pkg/gravatar"
)

func (r *queryResolver) Profile(ctx context.Context) (UserProfileResult, error) {
	if userCtx := auth.UserFromCtx(ctx); userCtx != nil {
		user, err := r.useCases.GetUserProfile(ctx, userCtx.ID)
		if err != nil {
			if errors.Is(err, crm.ErrForbidden) {
				return forbidden()
			}

			slog.ErrorContext(ctx, "can't get profile", slog.String("user", userCtx.ID), slog.Any("err", err))

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
