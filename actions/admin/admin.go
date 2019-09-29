package admin

import (
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/views/admin"
	"github.com/gobuffalo/buffalo"
)

type Routes struct {
	Manifest *assets.Manifest
}

func HomeRoute(manifest *assets.Manifest) buffalo.Handler {
	return func(c buffalo.Context) error {
		view, err := admin.Home(c, manifest)
		if err != nil {
			return err
		}
		return c.Render(200, render.EltToRenderer(view))
	}
}
