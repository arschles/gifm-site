package render

import (
	"io"
	"strings"
)

type Elt interface {
	ToHTML() (io.Reader, error)
}

type emptyElt struct{}

func (e emptyElt) ToHTML() (io.Reader, error) {
	r := strings.NewReader("")
	return r, nil
}
func EmptyElt() emptyElt {
	return emptyElt{}
}
