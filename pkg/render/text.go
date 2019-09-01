package render

import (
	"io"
	"strings"
)

func Text(s string) Elt {
	return text{t: s}
}

type text struct {
	t string
}

func (t text) ToHTML() (io.Reader, error) {
	return strings.NewReader(t.t), nil
}
