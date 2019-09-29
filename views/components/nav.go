package components

import (
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
)

// Nav is the navbar component
func Nav() render.Elt {
	return render.Tag("nav",
		render.TagOpts{"class": "navbar navbar-expand-lg navbar-light bg-light"},
		tags.Div(render.TagOpts{"class": "container-fluid"},
			render.Tag("a",
				render.TagOpts{"class": "navbar-brand", "href": "/"},
				render.Text("Go in 5 Minutes"),
			),
			render.Tag(
				"button",
				render.TagOpts{
					"class":         "navbar-toggler",
					"type":          "button",
					"data-toggle":   "collapse",
					"data-target":   "#navbarSupportedContent",
					"aria-controls": "navbarSupportedContent",
					"aria-expanded": "false",
					"aria-label":    "Toggle navigation",
				},
				render.Tag("span", render.TagOpts{"class": "navbar-toggler-icon"}),
			),
			render.Tag(
				"div",
				render.TagOpts{
					"class": "navbar-collapse collapse",
					"id":    "navbarSupportedContent",
				},
				linkList(),
				// render.Tag("ul", render.TagOpts{"class": "navbar-nav mr-auto"}, linkList()),
			),
		),
	)
}

func liNavItem(contents ...render.Elt) render.Elt {
	return render.Tag("li", render.TagOpts{"class": "nav-item pl-5"}, contents...)
}

func linkList() render.Elt {
	return render.Tag("ul", render.TagOpts{"class": "navbar-nav mr-auto"},
		liNavItem(tags.A("/screencasts", render.EmptyOpts(), "Episodes")),
		liNavItem(tags.A("/subscribe", render.EmptyOpts(), "Subscribe")),
		liNavItem(tags.A("/backers", render.EmptyOpts(), "Backers")),
		// TODO:
		// <span class="glyphicon glyphicon-new-window" aria-hidden="true"></span>
		// inside the <a>
		liNavItem(tags.A("https://bitly.com/goin5minutes", render.TagOpts{
			"target": "_blank",
		}, "Code on GitHub")),
		// TODO:
		// <span class="glyphicon glyphicon-new-window" aria-hidden="true"></span>
		// inside the <a>
		liNavItem(tags.A("https://www.patreon.com/goin5minutes", render.TagOpts{
			"target": "_blank",
		}, "Become a Patron")),
	)
}
