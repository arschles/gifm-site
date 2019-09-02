package views

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/components"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

func Home(manifest *assets.Manifest) (render.Elt, error) {
	ytSubscribeButton := tags.Div(
		render.TagOpts{"class": "col-sm-6 text-left"},
		tags.Div(
			render.TagOpts{
				"class":          "g-ytsubscribe",
				"data-channelid": "UC2GHqYE3fVJMncbrRd8AqcA",
				"data-layout":    "full",
				"data-count":     "default",
			},
		),
	)

	ghStarButton := tags.Div(
		render.TagOpts{"class": "col-sm-6 text-right"},
		render.Tag("iframe", render.TagOpts{
			"src":   "https://ghbtns.com/github-btn.html?user=arschles&repo=go-in-5-minutes&type=star&count=true&size=large",
			"style": "border:0; width:160px; height:30px; overflow:hidden",
		}),
	)

	return components.Base(
		manifest,
		tags.Div(render.TagOpts{"class": "container text-center"},
			tags.Div(
				render.TagOpts{"class": "row text-center"},
				tags.Div(render.TagOpts{"class": "jumbotron", "id": "gifm-landing-jumbotron"},
					tags.P(
						render.TagOpts{"class": "gifm-subtitle"},
						render.Text("Short Screencasts for Busy Gophers"),
					),
					render.Tag(
						"h1",
						render.TagOpts{"class": "gifm-title"},
						tags.Span(
							render.TagOpts{"id": "gifm-rotating-topics"},
							render.Text("Testing, Debugging, HTTP Servers, Websockets, Gorilla Mux, Benchmarking"),
						),
					),
					tags.P(render.EmptyOpts(),
						tags.A(
							"/screencasts",
							render.TagOpts{
								"class": "btn btn-success btn-lg gifm-start-button",
								"role":  "button",
							},
							"Start Watching Now",
						),
					),
				),
				tags.Div(render.TagOpts{"class": "row text-center"},
					tags.P(render.TagOpts{"class": "lead"},
						render.Text(`Go In 5 Minutes is a screencast series with the most concise, practical <a href="http://golang.org">Go</a> screencasts on the web.`),
					),
					tags.P(render.EmptyOpts(),
						tags.Small(
							render.EmptyOpts(),
							render.Text(`The series was featured on <a href="https://gotime.fm">GoTime</a>. Check out <a href="https://changelog.com/gotime/18">the episode</a>!`),
						),
					),
				),
				tags.Div(
					render.TagOpts{"class": "row text-center"},
					tags.Div(
						render.TagOpts{"class": "row text-center"},
						ghStarButton,
						ytSubscribeButton,
					),
				),
			),
		),
	)
}
