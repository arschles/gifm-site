package render
import (
	"bytes"
	"io"
	"strings"
	"fmt"

	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/tags"
)
type Elt render.Renderer

func Tag(name string, opts tags.Options, contents ...Elt) Elt {
	return elt{name: name, opts: opts, children: contents}
}

type elt struct {
	name string
	opts tags.Options
	children []Elt
}

func (e elt) ContentType() string {
	return "text/html"
}

func (e elt) Render(w io.Writer, d render.Data) error {
	children := make([]string, len(e.children))
	for i, child := range e.children {
		var childBytes bytes.Buffer
		if err := child.Render(&childBytes, d); err != nil {
			return err
		}
		children[i] = string(childBytes.Bytes())
	}
	toWrite := fmt.Sprintf(
		`<%s %s>
			%s
		</%s>`,
		e.name,
		optsAsString(e.opts),
		strings.Join(children, "\n"),
	)
	_, err := w.Write([]byte(toWrite))
	return err
}

func optsAsString(opts tags.Options) string {
	attrs := []string{}
	for k, v := range opts {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, v))
	}
	return strings.Join(attrs, " ")
}
