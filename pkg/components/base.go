package components

import (
	"io"
	"strings"

	"github.com/arschles/go-in-5-minutes-site/pkg/render"
)

type emptyElt struct{}

func (e emptyElt) ToHTML() (io.Reader, error) {
	return strings.NewReader(""), nil
}

type baseTag struct {
	baseElt render.Elt
}

func (b baseTag) ToHTML() (io.Reader, error) {
	preamble := strings.NewReader("<!doctype html>")
	remaining, err := b.baseElt.ToHTML()
	if err != nil {
		return nil, err
	}
	return io.MultiReader(preamble, remaining), nil
}

func Base(body render.Elt) render.Elt {
	return baseTag{
		baseElt: render.Tag(
			"html",
			render.TagOpts{"lang": "en"},
			head(),
			render.Tag("body", render.EmptyOpts(),
				nav(),
				body),
		),
	}
}
