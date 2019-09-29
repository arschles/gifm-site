package admin

import (
	"github.com/arschles/gifm-site/pkg/components/bootstrap"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
)

func nav() render.Elt {
	standardColOpts := render.TagOpts{"class": "col-sm-3"}
	return bootstrap.NewGrid(render.TagOpts{
		"class": "mt-4",
	}).WithRow(
		bootstrap.NewRow(render.EmptyOpts()).WithCol(
			bootstrap.NewCol(standardColOpts).WithChild(
				tags.A("/", render.EmptyOpts(), "Admin Home"),
			),
		).WithCol(
			bootstrap.NewCol(standardColOpts).WithChild(
				tags.A(
					"/admin/screencasts/new",
					render.EmptyOpts(),
					"Create a new screencast",
				),
			),
		),
	)
}
