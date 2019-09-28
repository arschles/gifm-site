package admin

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/components"
	bts "github.com/arschles/go-in-5-minutes-site/pkg/components/bootstrap"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
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
						render.NewTag("h2").WithChild(render.Text("Create a new Screencast")),
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
