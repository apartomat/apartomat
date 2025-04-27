package album

import (
	"encoding/xml"
	"fmt"
	"github.com/apartomat/apartomat/internal/store/albums"
)

const ns = "http://www.w3.org/2000/svg"

type Image struct {
	XMLName xml.Name `xml:"image"`

	Href   string `xml:"href,attr"`
	X      Length `xml:"x,attr"`
	Y      Length `xml:"y,attr"`
	Width  Length `xml:"width,attr"`
	Height Length `xml:"height,attr"`
}

type Length struct {
	Value float64
	Unit  Unit
}

type Unit string

const (
	mm Unit = "mm"
)

func (d Length) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: fmt.Sprintf("%.0f%s", d.Value, d.Unit),
	}, nil
}

type Size struct {
	Width  Length
	Height Length
}

var (
	defaultPageSize map[albums.PageSize]map[albums.PageOrientation]Size
)

func init() {
	defaultPageSize = map[albums.PageSize]map[albums.PageOrientation]Size{
		albums.A3: {
			albums.Portrait: {
				Width: Length{
					Value: 297,
					Unit:  mm,
				},
				Height: Length{
					Value: 420,
					Unit:  mm,
				},
			},
			albums.Landscape: {
				Width: Length{
					Value: 420,
					Unit:  mm,
				},
				Height: Length{
					Value: 297,
					Unit:  mm,
				},
			},
		},
		albums.A4: {
			albums.Portrait: {
				Width: Length{
					Value: 210,
					Unit:  mm,
				},
				Height: Length{
					Value: 297,
					Unit:  mm,
				},
			},
			albums.Landscape: {
				Width: Length{
					Value: 297,
					Unit:  mm,
				},
				Height: Length{
					Value: 210,
					Unit:  mm,
				},
			},
		},
	}
}
