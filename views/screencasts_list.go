package views

import (
	"strconv"

	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
	"github.com/arschles/gifm-site/views/components/bootstrap"
)

// ScreencastsList returns a slice of rows, one for each screencast in screencasts.
//
// isAdmin will be used to customize the view of each row
func ScreencastsList(
	screencasts *models.Screencasts,
	isAdmin bool,
) []bootstrap.Row {
	rows := make([]bootstrap.Row, len(*screencasts))
	stdColOpts := render.TagOpts{"class": "col-sm text-left"}

	stdScreencastRow := bootstrap.NewRow(render.TagOpts{
		"class": "mt-3 shadow-sm p-3 mb-5 bg-white rounded",
	})
	for i, screencast := range *screencasts {
		var editButton render.Elt
		editButton = render.EmptyElt()
		if isAdmin {
			editButton = render.NewTag("a").
				WithOpts(render.TagOpts{
					"class": "btn btn-outline-primary mt-1",
					"type":  "button",
					"role":  "button",
					// TODO: this link doesn't work!
					"href": "/admin/screencasts/" + screencast.ID.String() + "/edit",
				}).
				WithChild(
					render.Text("Edit"),
				)
		}
		rows[i] =
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
					editButton,
				),
			)
	}

	return rows
}
