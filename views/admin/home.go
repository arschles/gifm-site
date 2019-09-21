package admin

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/components"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
)

func Home(manifest *assets.Manifest) (render.Elt, error) {
	return components.Base(manifest, render.Text("Hello admin!"))
}
