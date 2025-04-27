package album

import (
	"encoding/xml"
	"fmt"
	"github.com/apartomat/apartomat/internal/store/albums"
)

type visualizationSvg struct {
	XMLName xml.Name `xml:"svg"`

	NS string `xml:"xmlns,attr"`
	ID string `xml:"id,attr"`

	Width    Length `xml:"width,attr"`
	Height   Length `xml:"height,attr"`
	Overflow string `xml:"overflow,attr"`

	Visualization visualization
}

type visualization struct {
	XMLName xml.Name `xml:"image"`

	Href   string `xml:"href,attr"`
	X      Length `xml:"x,attr"`
	Y      Length `xml:"y,attr"`
	Width  Length `xml:"width,attr"`
	Height Length `xml:"height,attr"`
}

func Visualization(format albums.PageSize, orientation albums.PageOrientation) func(id int, path string) (string, error) {
	var (
		defaultMargin = 0.5 * 25.4
		bindingMargin = 25.4

		x, y float64

		width  = defaultPageSize[format][orientation].Width.Value
		height = defaultPageSize[format][orientation].Height.Value
	)

	switch orientation {
	case albums.Portrait:
		x = defaultMargin
		y = bindingMargin

		width -= 2 * defaultMargin
		height -= bindingMargin + defaultMargin
	case albums.Landscape:
		x = bindingMargin
		y = defaultMargin

		width -= bindingMargin + defaultMargin
		height -= 2 * defaultMargin
	}

	return func(id int, path string) (string, error) {
		svg := visualizationSvg{
			NS:       ns,
			ID:       fmt.Sprintf("id-%d", id),
			Overflow: "visible",
			Width:    defaultPageSize[format][orientation].Width,
			Height:   defaultPageSize[format][orientation].Height,
			Visualization: visualization{
				Href: path,
				X: Length{
					Value: x,
					Unit:  mm,
				},
				Y: Length{
					Value: y,
					Unit:  mm,
				},
				Width: Length{
					Value: width,
					Unit:  mm,
				},
				Height: Length{
					Value: height,
					Unit:  mm,
				},
			},
		}

		b, err := xml.Marshal(svg)
		if err != nil {
			return "", err
		}

		return string(b), nil
	}
}
