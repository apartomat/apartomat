package graphql

import (
	"context"
)

func (r *rootResolver) AlbumPageCover() AlbumPageCoverResolver {
	return &albumPageCoverResolver{r}
}

type albumPageCoverResolver struct {
	*rootResolver
}

func (r *albumPageCoverResolver) SVG(ctx context.Context, obj *AlbumPageCover) (AlbumPageSVGResult, error) {
	return SVG{
		SVG: `<svg width="420mm" height="297mm" xmlns="http://www.w3.org/2000/svg" overflow="visible"><rect x="0" y="0" width="100%" height="100%" fill="lightgray"></rect><rect x="30mm" y="10mm" width="380mm" height="277mm" fill="white"></rect><text color="black" x="40mm" y="30mm" font-size="56px" font-family="Arial, Helvetica, sans-serif">PUHOVA</text><text color="black" x="40mm" y="50mm" font-size="36px" font-family="Arial, Helvetica, sans-serif">Новосибирск 2024</text><text color="black" x="40mm" y="70mm" font-size="24px" font-family="Arial, Helvetica, sans-serif">Дизайн-проект интерьера квартиры 155,87 м²</text><text color="black" x="40mm" y="90mm" font-size="24px" font-family="Arial, Helvetica, sans-serif">Зыряновская, 51</text></svg>`,
	}, nil
}

func (r *albumPageCoverResolver) Cover(
	ctx context.Context,
	obj *AlbumPageCover,
) (AlbumPageCoverResult, error) {
	var (
		cov interface{} = obj.Cover
	)

	switch c := cov.(type) {
	case *CoverUploaded:
		return c, nil
	case *Cover:
		return notImplementedYetError()
	default:
		r.logger.Error("unknown obj type")
		return serverError()
	}
}
