package views

import (
	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/security"
	"github.com/arschles/gifm-site/pkg/tags"
	"github.com/arschles/gifm-site/views/components"
	"github.com/arschles/gifm-site/views/components/bootstrap"
	"github.com/gobuffalo/buffalo"
)

func Screencasts(
	c buffalo.Context,
	manifest *assets.Manifest,
	screencasts *models.Screencasts,
) (render.Elt, error) {
	stdColOpts := render.TagOpts{"class": "col-sm text-left"}
	header := render.Tag(
		"h1",
		render.TagOpts{"class": "page-title"},
		render.Text("All Screencasts"),
		tags.Div(
			render.TagOpts{
				"class": "d-inline-flex p-2 bd-highlight",
			},
			render.NewTag("small").WithOpt("class", "text-muted").WithChild(
				render.Text("Newest on top"),
			),
		),
	)
	screencastRows := ScreencastsList(screencasts, security.IsAdmin(c))
	grid := bootstrap.NewGrid(render.EmptyOpts()).WithRow(
		bootstrap.NewRow(render.EmptyOpts()).WithCol(
			bootstrap.NewCol(stdColOpts).WithChild(header),
		),
	).WithRows(screencastRows...)

	return components.Base(
		c,
		manifest,
		grid,
	)
}
