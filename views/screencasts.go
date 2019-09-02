package views

import (
	"github.com/arschles/go-in-5-minutes-site/models"
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/components"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

func Screencasts(manifest *assets.Manifest, screencasts *models.Screencasts) (render.Elt, error) {
	screencastsList := render.NewTag("ul")
	for _, screencast := range *screencasts {
		screencastsList = screencastsList.WithChild(
			render.Tag("li", render.TagOpts{}, render.Text(screencast.Title)),
		)
	}

	return components.Base(
		manifest,
		tags.Div(render.TagOpts{"class": "container text-center"},
			tags.Div(render.TagOpts{"class": "row"},
				render.Tag(
					"h1",
					render.TagOpts{"class": "page-title"},
					render.Text("All Screencasts"),
					tags.Div(
						render.TagOpts{
							"class": "d-inline-flex p-2 bd-highlight",
						},
						render.Text("Newest on top"),
					),
				),
			),
			tags.Div(render.TagOpts{}, screencastsList),
		),
	)
}
