package render

import (
	"io"

	"github.com/gobuffalo/buffalo/render"
)
func Text(s string) Elt {
	return text{t: s}
}

type text struct {
	t string
}

func (t text) ContentType() string {
	return "text/plain"
}

func (t text) Render(w io.Writer, _ render.Data) error {
	_, err := w.Write([]byte(t.t))
	return err
}