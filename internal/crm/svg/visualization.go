package svg

import (
	"bytes"
	"html/template"
)

const visualizationTpl = `
<svg id="svg-{{ .ID }}" width="420mm" height="297mm" xmlns="http://www.w3.org/2000/svg" overflow="visible">
	<image href="{{ .Path }}" x="30mm" y="10mm" width="380mm" height="277mm"/>
</svg>
`

var (
	visualization *template.Template
)

func init() {
	t, err := template.New("visualization").Parse(visualizationTpl)
	if err != nil {
		panic(err)
	}

	visualization = t
}

func Visualization(id int, path string) (string, error) {
	var (
		b = new(bytes.Buffer)
	)

	if err := visualization.Execute(b, struct {
		ID   int
		Path string
	}{id, path}); err != nil {
		return "", err
	}

	return b.String(), nil
}
