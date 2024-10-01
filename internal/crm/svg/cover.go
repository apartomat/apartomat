package svg

import (
	"bytes"
	"html/template"
)

const uploadedCoverTpl = `
<svg id="svg-{{ .ID }}" width="420mm" height="297mm" xmlns="http://www.w3.org/2000/svg" overflow="visible">
	<image href="{{ .Path }}" x="0" y="0" width="420mm" height="297mm"/>
</svg>
`

var (
	uploadedCover *template.Template
)

func init() {
	t, err := template.New("visualization").Parse(uploadedCoverTpl)
	if err != nil {
		panic(err)
	}

	uploadedCover = t
}

func UploadedCover(id int, path string) (string, error) {
	var (
		b = new(bytes.Buffer)
	)

	if err := uploadedCover.Execute(b, struct {
		ID   int
		Path string
	}{id, path}); err != nil {
		return "", err
	}

	return b.String(), nil
}

func Cover() string {
	return `<svg width="420mm" height="297mm" xmlns="http://www.w3.org/2000/svg" overflow="visible"><rect x="0" y="0" width="100%" height="100%" fill="lightgray"></rect><rect x="30mm" y="10mm" width="380mm" height="277mm" fill="white"></rect><text color="black" x="40mm" y="30mm" font-size="56px" font-family="Arial, Helvetica, sans-serif">PUHOVA</text><text color="black" x="40mm" y="50mm" font-size="36px" font-family="Arial, Helvetica, sans-serif">Новосибирск 2024</text><text color="black" x="40mm" y="70mm" font-size="24px" font-family="Arial, Helvetica, sans-serif">Дизайн-проект интерьера квартиры 155,87 м²</text><text color="black" x="40mm" y="90mm" font-size="24px" font-family="Arial, Helvetica, sans-serif">Зыряновская, 51</text></svg>`
}
