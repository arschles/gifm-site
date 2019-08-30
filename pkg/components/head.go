package components
import (
	"github.com/gobuffalo/tags"
	rndr "github.com/arschles/go-in-5-minutes-site/pkg/render"
)
func meta(opts tags.Options) rndr.Elt {
	return rndr.Tag("meta", opts, emptyElt{})
}

func head() rndr.Elt {
	return rndr.Tag("head", emptyOpts(), 
		meta(tags.Options{
			"name": "viewport",
			"content": "width=device-width, initial-scale=1",
		}),
		meta(tags.Options{"charset": "utf-8"}),
		meta(tags.Options{
			"name": "csrf-param",
			"content": "authenticity-token",
		}),
		meta(tags.Options{
			"name": "csrf-token",
			// TODO: make random! see https://github.com/gobuffalo/mw-csrf/blob/master/csrf.go
			"authenticity_token": "abvsfasdf",
		}),
		rndr.Tag("link", tags.Options{
			"rel": "icon",
			"href": "images/favicon.ico",
		}, emptyElt{}),
		// TODO: CSS and JS
		rndr.Tag("title", emptyOpts(), rndr.Text("Go in 5 Minutes")),
	)
}