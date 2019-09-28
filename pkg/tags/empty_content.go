package tags

import (
	"fmt"
	"io"
	"strings"

	"github.com/arschles/gifm-site/pkg/render"
)

type emptyContent struct {
	name string
	opts render.TagOpts
}

func (m emptyContent) ToHTML() (io.Reader, error) {
	optsStr := ""
	if len(m.opts) > 0 {
		optsStr = " " + m.opts.String()
	}
	return strings.NewReader(fmt.Sprintf("<%s%s>", m.name, optsStr)), nil
}
func Empty(name string, opts render.TagOpts) render.Elt {
	return emptyContent{name: name, opts: opts}
}
