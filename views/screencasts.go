package views

import (
	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/components"
	"github.com/arschles/gifm-site/pkg/components/bootstrap"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
)

func Screencasts(
	authenticityToken string,
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
			render.Tag(
				"mark",
				render.TagOpts{"class": "small"},
				render.Text("newest on top"),
			),
		),
	)
	grid := bootstrap.NewGrid(render.EmptyOpts()).WithRow(
		bootstrap.NewRow(render.EmptyOpts()).WithCol(
			bootstrap.NewCol(stdColOpts).WithChild(header),
		),
	)

	stdScreencastRow := bootstrap.NewRow(render.TagOpts{
		"class": "mt-3 shadow-sm p-3 mb-5 bg-white rounded",
	})
	for _, screencast := range *screencasts {
		grid = grid.WithRow(
			stdScreencastRow.WithCol(
				bootstrap.NewCol(stdColOpts).WithChildren(
					render.NewTag("h2").WithChild(
						render.Text(screencast.Title),
					),
					render.NewTag("small").WithChildren(
						render.Text(screencast.Intro),
					),
					// tags.A("/screencasts/"+screencast.ID.String(), render.EmptyOpts(), "details"),
					// <button type="button" class="btn btn-primary">Primary</button>
					render.NewTag("div").WithOpt("class", "mt-1").WithChild(
						tags.A("/screencasts/"+screencast.ID.String(), render.TagOpts{
							"class": "btn btn-outline-primary",
							"type":  "button",
							"role":  "button",
						}, "Details"),
					),
				),
			),
		)
	}

	return components.Base(
		authenticityToken,
		manifest,
		grid,
	)
}
