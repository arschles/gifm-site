package render

import (
	"io"

	"github.com/gobuffalo/buffalo/render"
)

func EltToRenderer(elt Elt) render.Renderer {
	return eltToRenderer{elt}
}

type eltToRenderer struct {
	e Elt
}

func (e eltToRenderer) ContentType() string { return "text/html" }
func (e eltToRenderer) Render(w io.Writer, _ render.Data) error {
	reader, err := e.e.ToHTML()
	if err != nil {
		return err
	}
	_, err = io.Copy(w, reader)
	return err
}
