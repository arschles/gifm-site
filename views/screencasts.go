package views

import (
	"github.com/arschles/go-in-5-minutes-site/models"
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/components"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

func Screencasts(authenticityToken string, manifest *assets.Manifest, screencasts *models.Screencasts) (render.Elt, error) {
	screencastsList := render.NewTag("div").WithOpt("class", "container-fluid")
	for _, screencast := range *screencasts {
		screencastsList = screencastsList.WithChild(Screencast(screencast))
	}

	return components.Base(
		authenticityToken,
		manifest,
		tags.Div(render.TagOpts{"class": "container text-center"},
			tags.Div(render.TagOpts{"class": "row text-center"},
				tags.Div(render.TagOpts{"class": "col-sm text-center"},
					render.Tag(
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
					),
				),
			),
			screencastsList,
		),
	)
}
