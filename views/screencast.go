package views

import (
	"github.com/arschles/go-in-5-minutes-site/models"
	"github.com/arschles/go-in-5-minutes-site/pkg/components/bootstrap"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

// Screencast renders a single screencast
func Screencast(screencast models.Screencast) render.Elt {
	return tags.Div(render.TagOpts{"class": "row", "style": "margin-top:2em;"},
		tags.Div(render.TagOpts{"class": "col-sm-12"},
			bootstrap.Card(screencast.Title, screencast.Intro),
		),
	)
}
