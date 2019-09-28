package admin

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/views/admin"
	"github.com/gobuffalo/buffalo"
)

type Routes struct {
	Manifest *assets.Manifest
}

func HomeRoute(manifest *assets.Manifest) buffalo.Handler {
	return func(c buffalo.Context) error {
		view, err := admin.Home(manifest)
		if err != nil {
			return err
		}
		return c.Render(200, render.EltToRenderer(view))
	}
}
