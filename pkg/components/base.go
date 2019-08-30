package components
import (
	"io"

	"github.com/gobuffalo/tags"
	"github.com/gobuffalo/buffalo/render"
	rndr "github.com/arschles/go-in-5-minutes-site/pkg/render"
)

func emptyData() render.Data { return map[string]interface{}{} }
func emptyOpts() tags.Options { return map[string]interface{}{} }
type emptyElt struct{}
func (e emptyElt) ContentType() string {return ""}
func (e emptyElt) Render(io.Writer, render.Data) error {
	return nil
}

func Base(body rndr.Elt) rndr.Elt {
	return rndr.Tag(
		"html",
		emptyOpts(),
		head(),
		body,
	)
}