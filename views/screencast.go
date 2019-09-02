package views

import (
	"github.com/arschles/go-in-5-minutes-site/models"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

func Screencast(screencast models.Screencast) render.Elt {
	return tags.Div(render.TagOpts{"class": "row", "style": "margin-top:2em;"},
		tags.Div(render.TagOpts{"class": "col-sm-12"},
			tags.Div(
				render.TagOpts{"class": "card"},
				tags.Div(
					render.TagOpts{
						"class": "card-body",
					},
					render.Tag("h5", render.TagOpts{}, render.Text(screencast.Title)),
					tags.P(render.TagOpts{"class": "card-text"}, render.Text(screencast.Intro)),
				),
			),
		),
	)
}
