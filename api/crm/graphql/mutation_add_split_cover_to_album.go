package graphql

import (
	"context"
	"errors"
	"log/slog"

	"github.com/apartomat/apartomat/internal/crm"
)

func (r *mutationResolver) AddSplitCoverToAlbum(
	ctx context.Context,
	albumID string,
	input AddSplitCoverToAlbumInput,
) (AddSplitCoverToAlbumResult, error) {
	withQRValue := false
	if input.WithQR != nil {
		withQRValue = *input.WithQR
	}

	_, err := r.crm.AddSplitCoverToAlbum(
		ctx,
		albumID,
		crm.SplitCoverForAddToAlbum{
			Title:     input.Title,
			Subtitle:  input.Subtitle,
			ImgFileID: input.ImgFileID,
			WithQR:    withQRValue,
			City:      input.City,
			Year:      input.Year,
		},
	)
	if err != nil {
		if errors.Is(err, crm.ErrForbidden) {
			return forbidden()
		}
		if errors.Is(err, crm.ErrNotFound) {
			return notFound()
		}

		slog.ErrorContext(
			ctx,
			"can't add split cover to album",
			slog.String("album", albumID),
			slog.Any("err", err),
		)

		return serverError()
	}

	splitCover := &SplitCover{
		Title:    input.Title,
		Subtitle: input.Subtitle,
		Image: &File{
			ID: input.ImgFileID,
		},
		QRCodeSrc: nil,
		City:      input.City,
		Year:      input.Year,
		Variant:   SplitCoverVariantImageOnTheRight,
	}

	return SplitCoverAdded{Cover: splitCover}, nil
}
