package admin

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/components"
	bts "github.com/arschles/go-in-5-minutes-site/pkg/components/bootstrap"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
)

func Home(manifest *assets.Manifest) (render.Elt, error) {
	return components.Base(
		manifest,
		render.NewTag("div").WithOpt("class", "container").WithChild(
			nav(),
		).WithChild(
			bts.NewGrid(render.TagOpts{
				"class": "mt-2",
			}).WithRow(
				bts.NewRow(render.EmptyOpts()).WithCol(
					bts.NewCol(render.TagOpts{
						"class": "col-sm text-center",
					}).WithChild(
						render.NewTag("h2").WithChild(render.Text("Create a new Screencast")),
					),
				),
			).WithRow(
				bts.NewRow(render.EmptyOpts()).WithCol(
					bts.NewCol(render.TagOpts{
						"class": "col-sm text-center",
					}).WithChild(
						ScreencastForm(),
					),
				),
			),
		),
	)
}
