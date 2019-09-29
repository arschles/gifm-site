package views

import (
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/views/components"
	"github.com/arschles/gifm-site/views/components/bootstrap"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
	"github.com/gobuffalo/buffalo"
)

func Home(c buffalo.Context, manifest *assets.Manifest) (render.Elt, error) {

	ytSubscribeButton := tags.Div(
		render.TagOpts{"class": "col text-left"},
		render.Tag("script", render.TagOpts{"src": "https://apis.google.com/js/platform.js"}),
		tags.Div(
			render.TagOpts{
				"class":          "g-ytsubscribe",
				"data-channelid": "UC2GHqYE3fVJMncbrRd8AqcA",
				"data-layout":    "default",
				"data-count":     "default",
			},
		),
	)

	ghStarButton := tags.Div(
		render.TagOpts{"class": "col text-right"},
		render.Tag("iframe", render.TagOpts{
			"src":   "https://ghbtns.com/github-btn.html?user=arschles&repo=go-in-5-minutes&type=star&count=true&size=large",
			"style": "border:0; width:160px; height:30px; overflow:hidden",
		}),
	)

	return components.Base(
		c,
		manifest,
		tags.Div(render.TagOpts{"class": "container text-center mt-5"},
			tags.Div(
				render.TagOpts{"class": "row text-center mt-5"},
				render.NewTag("div").WithOpt("class", "col-sm").WithChild(
					render.NewTag("h1").WithOpt("class", "text-monospace").WithText("Short Screencasts for Busy Gophers"),
				),
			),
			render.NewTag("div").WithOpt("class", "row text-center mt-3").WithChild(
				render.NewTag("div").WithOpt("class", "col-sm").WithChild(
					tags.A(
						"/screencasts",
						render.TagOpts{
							"class": "btn btn-link btn-lg",
							"role":  "button",
							"id":    "watch-now-button",
						},
						"Start Watching Now",
					),
				),
			),
			render.NewTag("div").WithOpt("class", "row text-center mt-5").WithChildren(
				render.NewTag("div").WithOpt("class", "col-sm").WithChild(
					bootstrap.Card("<mark>Testing</mark>", "Advanced practices for testing even complicated codebases"),
				),
				render.NewTag("div").WithOpt("class", "col-sm").WithChild(
					bootstrap.Card("<mark>Debugging</mark>", "Debugging real-life apps, even with lots of concurrency"),
				),
				render.NewTag("div").WithOpt("class", "col-sm").WithChild(
					bootstrap.Card("<mark>HTTP Servers</mark>", "Writing awesome, full-featured web applications"),
				),
				render.NewTag("div").WithOpt("class", "col-sm").WithChild(
					bootstrap.Card("<mark>Websockets</mark>", "Integrating real-time features into your web applications"),
				),
			),
			render.NewTag("div").WithOpt("class", "row text-center mt-5").WithChild(
				render.NewTag("div").WithOpt("class", "col-sm text-center").WithChild(
					render.NewTag("p").WithOpt("class", "lead").WithText(
						`Go In 5 Minutes is a screencast series with the most concise, practical <a href="http://golang.org">Go</a> screencasts on the web.`,
					),
				).WithChild(
					render.NewTag("p").WithChild(
						render.NewTag("small").WithText(
							`The series was featured on <a href="https://gotime.fm">GoTime</a>. Check out <a href="https://changelog.com/gotime/18">the episode</a>!`,
						),
					),
				),
			),
			tags.Div(
				render.TagOpts{"class": "row"},
				ghStarButton,
				ytSubscribeButton,
			),
		),
	)
}
