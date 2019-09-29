package admin

import (
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/components"
	bts "github.com/arschles/gifm-site/pkg/components/bootstrap"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/pkg/tags"
)

// Home returns the homepage for the admin tool
func Home(manifest *assets.Manifest, authenticityToken string) (render.Elt, error) {
	standardColOpts := render.TagOpts{
		"class": "col-sm text-center",
	}
	return components.Base(
		authenticityToken,
		manifest,
		render.NewTag("div").WithOpt("class", "container").WithChild(
			nav(),
		).WithChild(
			bts.NewGrid(standardColOpts).WithRow(
				bts.NewRow(render.EmptyOpts()).WithCol(
					bts.NewCol(render.TagOpts{
						"class": "col-sm text-center",
					}).WithChild(
						render.NewTag("h2").WithChild(render.Text("Admin Home")),
					),
				),
			).WithRow(
				bts.NewRow(render.EmptyOpts()).WithCol(
					bts.NewCol(standardColOpts).WithChild(
						tags.A(
							"/admin/screencasts/new",
							render.EmptyOpts(),
							"Create a new screencast",
						),
					),
				),
			),
			// ).WithRow(
			// 	bts.NewRow(render.EmptyOpts()).WithCol(
			// 		bts.NewCol(render.TagOpts{
			// 			"class": "col-sm text-center",
			// 		}).WithChild(
			// 			ScreencastForm(authenticityToken, "/admin/screencasts"),
			// 		),
			// 	),
			// ),
		),
	)
}
