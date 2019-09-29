package admin

import (
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
	"github.com/arschles/gifm-site/views/components/bootstrap"
)

func nav() render.Elt {
	stdCol := bootstrap.NewCol(render.TagOpts{"class": "col-sm"})
	return bootstrap.NewGrid(render.TagOpts{
		"class": "mt-4",
	}).WithRow(
		bootstrap.NewRow(render.EmptyOpts()).WithCol(
			stdCol.WithChild(
				tags.A("/admin", render.EmptyOpts(), "Admin Home"),
			),
		).WithCol(
			stdCol.WithChild(
				tags.A(
					"/admin/screencasts/new",
					render.EmptyOpts(),
					"Create a new screencast",
				),
			),
		),
	)
}
