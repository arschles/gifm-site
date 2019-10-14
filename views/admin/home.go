package admin

import (
	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/views"
	"github.com/arschles/gifm-site/views/components"
	bts "github.com/arschles/gifm-site/views/components/bootstrap"
	"github.com/gobuffalo/buffalo"
)

// Home returns the homepage for the admin tool
func Home(c buffalo.Context, manifest *assets.Manifest, screencasts *models.Screencasts) (render.Elt, error) {
	standardColOpts := render.TagOpts{
		"class": "col-sm text-center",
	}
	screencastRows := views.ScreencastsList(screencasts, true)

	return components.Base(
		c,
		manifest,
		render.NewTag("div").WithOpt("class", "container").WithChild(
			bts.NewGrid(standardColOpts).WithRow(
				bts.NewRow(render.EmptyOpts()).WithCol(
					bts.NewCol(render.TagOpts{
						"class": "col-sm text-center mt-4",
					}).WithChild(
						render.NewTag("h2").WithChild(render.Text("Admin Home")),
					),
				),
			).WithRow(
				bts.NewRow(render.EmptyOpts()).WithCol(
					bts.NewCol(render.EmptyOpts()).WithChild(nav()),
				),
			).WithRows(screencastRows...),
		),
	)
}
