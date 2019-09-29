package views

import (
	"strconv"

	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/render"
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
					).WithChild(
						render.NewTag("small").WithOpt("class", "text-muted").WithChild(
							render.Text("Episode "+strconv.Itoa(screencast.EpisodeNum)),
						),
					),
					render.NewTag("h4").WithChildren(
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
		c,
		manifest,
		grid,
	)
}
