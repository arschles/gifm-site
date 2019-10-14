package admin

import (
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
)

func nav() render.Elt {
	links := []render.Elt{
		tags.A("/admin", render.EmptyOpts(), "Admin Home"),
		tags.A(
			"/admin/screencasts/new",
			render.TagOpts{"class": "ml-4"},
			"Create a new screencast",
		),
	}
	return tags.Div(render.EmptyOpts(), links...)
}
