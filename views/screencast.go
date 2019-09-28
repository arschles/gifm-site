package views

import (
	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/components/bootstrap"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
)

// Screencast renders a single screencast
func Screencast(screencast models.Screencast) render.Elt {
	return tags.Div(render.TagOpts{"class": "row", "style": "margin-top:2em;"},
		tags.Div(render.TagOpts{"class": "col-sm-12"},
			bootstrap.Card(screencast.Title, screencast.Intro),
		),
	)
}
