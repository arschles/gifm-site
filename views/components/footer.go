package components

import (
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
)

func footer() render.Elt {
	return render.NewTag("footer").WithChild(
		render.NewTag("div").WithOpt("class", "container mt-5 border-top").WithChild(
			render.NewTag("div").WithOpt("class", "row").WithChildren(
				render.NewTag("div").WithOpt("class", "col-md-6 text-left mt-3 text-secondary small").WithChild(
					tags.A("/subscribe", render.EmptyOpts(), "Sign up for the newsletter"),
				),
				render.NewTag("div").
					WithOpt("class", "col-md-6 text-right mt-3 text-secondary small").
					WithChildren(
						tags.P(
							render.EmptyOpts(),
							render.Text(
								`Glyphicons courtesy of <a href="http://glyphicons.com">glyphicons.com</a>`,
							),
						),
						tags.P(
							render.EmptyOpts(),
							render.Text(
								`Gopher icons attributed to <a href="http://reneefrench.blogspot.com/">Renee French</a>`,
							),
						),
						tags.P(
							render.EmptyOpts(),
							render.Text(
								`Proudly created with <a href="http://gobuffalo.io/">Buffalo</a>`,
							),
						),
					),
			),
		),
	)
}
