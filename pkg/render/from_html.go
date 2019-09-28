package render

import (
	"bytes"
	"io"

	"github.com/gobuffalo/buffalo/render"
)

type htmlRenderElt struct {
	contents []byte
}

func (e htmlRenderElt) ToHTML() (io.Reader, error) {
	return bytes.NewReader(e.contents), nil
}

// FromHTML brings an HTML file into a render.Elt
func FromHTML(tplName string, r *render.Engine, data map[string]interface{}) (Elt, error) {
	renderer := r.HTML(tplName)
	var b bytes.Buffer
	if err := renderer.Render(&b, data); err != nil {
		return nil, err
	}
	return htmlRenderElt{contents: b.Bytes()}, nil
}
