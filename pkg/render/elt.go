package render

import (
	"io"
)

type Elt interface {
	ToHTML() (io.Reader, error)
}
