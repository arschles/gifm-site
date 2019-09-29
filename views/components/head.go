package components

import (
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/security"
	"github.com/arschles/gifm-site/pkg/tags"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

func head(
	c buffalo.Context,
	manifest *assets.Manifest,
) (render.Elt, error) {
	authenticityToken := security.AuthenticityTokenFromCtx(c)
	jsElt, err := tags.JS(manifest, "application.js")
	if err != nil {
		return nil, errors.WithMessage(err, "Trying to construct <head>")
	}
	cssElt, err := tags.CSS(manifest, "application.css")
	if err != nil {
		return nil, errors.WithMessage(err, "Trying to construct <head>")
	}
	return render.Tag("head", render.EmptyOpts(),
		tags.Meta(render.TagOpts{
			"name":    "viewport",
			"content": "width=device-width, initial-scale=1, shrink-to-fit=no",
		}),
		tags.Meta(render.TagOpts{"charset": "utf-8"}),
		tags.Meta(render.TagOpts{
			"name":    "csrf-param",
			"content": security.TokenFieldName,
		}),
		tags.Meta(render.TagOpts{
			"name": "csrf-token",
			// TODO: make random! see https://github.com/gobuffalo/mw-csrf/blob/master/csrf.go
			"authenticity_token": authenticityToken,
		}),
		tags.Link(render.TagOpts{"rel": "icon", "href": "images/favicon.ico"}),
		jsElt,
		cssElt,
		render.Tag("title", render.EmptyOpts(), render.Text("Go in 5 Minutes")),
		tags.Link(render.TagOpts{
			"href": "//fonts.googleapis.com/css?family=Varela+Round|Cousine:400,700",
			"rel":  "stylesheet",
			"type": "text/css",
		}),
	), nil
}
