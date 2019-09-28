package actions

import (
	"fmt"
	"os"

	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/helpers"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr/v2"
)

var r *render.Engine
var assetsBox = packr.New("app:assets", "../public")
var parsedManifest *assets.Manifest

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.plush.html",

		// Box containing all of the templates:
		TemplatesBox: packr.New("app:templates", "../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// for non-bootstrap form helpers uncomment the lines
			// below and import "github.com/gobuffalo/helpers/forms"
			// forms.FormKey:     forms.Form,
			// forms.FormForKey:  forms.FormFor,
			"tag":       helpers.Tag,
			"div":       helpers.Div,
			"container": helpers.Container,
			"row":       helpers.Row,
			"text":      helpers.Text,
		},
	})
	goEnv := os.Getenv("GO_ENV")
	var err error
	parsedManifest, err = assets.ParseManifest(goEnv != "production", assetsBox)
	if err != nil {
		panic(fmt.Sprintf("Could not parse assets manifest: %s", err))
	}
	if goEnv != "production" {
		fmt.Println("Parsed manifest", *parsedManifest)
	}
}
