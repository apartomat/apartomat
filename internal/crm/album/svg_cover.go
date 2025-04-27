package album

import (
	"encoding/xml"
	"fmt"
	"github.com/apartomat/apartomat/internal/store/albums"
)

type coverSvg struct {
	XMLName xml.Name `xml:"svg"`

	NS string `xml:"xmlns,attr"`
	ID string `xml:"id,attr"`

	Width    Length `xml:"width,attr"`
	Height   Length `xml:"height,attr"`
	Overflow string `xml:"overflow,attr"`

	Cover Image
}

func UploadedCover(format albums.PageSize, orientation albums.PageOrientation) func(id int, path string) (string, error) {
	return func(id int, path string) (string, error) {
		svg := coverSvg{
			NS:       ns,
			ID:       fmt.Sprintf("id-%d", id),
			Overflow: "visible",
			Width:    defaultPageSize[format][orientation].Width,
			Height:   defaultPageSize[format][orientation].Height,
			Cover: Image{
				Href: path,
				X: Length{
					Value: 0,
					Unit:  mm,
				},
				Y: Length{
					Value: 0,
					Unit:  mm,
				},
				Width:  defaultPageSize[format][orientation].Width,
				Height: defaultPageSize[format][orientation].Height,
			},
		}

		b, err := xml.Marshal(svg)
		if err != nil {
			return "", err
		}

		return string(b), nil
	}
}

func Cover() string {
	return `<svg width="420mm" height="297mm" xmlns="http://www.w3.org/2000/svg" overflow="visible"><rect x="0" y="0" width="100%" height="100%" fill="lightgray"></rect><rect x="30mm" y="10mm" width="380mm" height="277mm" fill="white"></rect><text color="black" x="40mm" y="30mm" font-size="56px" font-family="Arial, Helvetica, sans-serif">PUHOVA</text><text color="black" x="40mm" y="50mm" font-size="36px" font-family="Arial, Helvetica, sans-serif">Новосибирск 2024</text><text color="black" x="40mm" y="70mm" font-size="24px" font-family="Arial, Helvetica, sans-serif">Дизайн-проект интерьера квартиры 155,87 м²</text><text color="black" x="40mm" y="90mm" font-size="24px" font-family="Arial, Helvetica, sans-serif">Зыряновская, 51</text></svg>`
}
