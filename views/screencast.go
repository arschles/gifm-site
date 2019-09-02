package views

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

func Screencast() render.Elt {
	return tags.Div(
		render.TagOpts{},
		render.Text("This is where screencasts go!"),
	)
}
