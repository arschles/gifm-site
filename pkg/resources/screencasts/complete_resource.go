package screencasts

import (
	"fmt"

	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/resources"
	"github.com/arschles/gifm-site/pkg/youtube"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	rndr "github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop"
)

// Resource is a complete buffalo.Resource for screencasts
type Resource struct {
	ReadOnlyResource
	resources.Base
	r        *render.Engine
	ytConfig youtube.Config
}

// NewResource creates a new Resource for screencasts
func NewResource(
	basePath string,
	r *rndr.Engine,
	manifest *assets.Manifest,
	ytConfig youtube.Config,
) buffalo.Resource {
	return Resource{
		ReadOnlyResource: NewReadOnlyResource(manifest),
		Base:             resources.NewBase(basePath),
		r:                r,
		ytConfig:         ytConfig,
	}
}

// New renders the form for creating a new Screencast.
// This function is mapped to the path GET /screencasts/new
func (r Resource) New(c buffalo.Context) error {
	// authenticityToken := forms.AuthenticityToken(c)
	// screencastForm := admin.ScreencastForm(authenticityToken, r.BasePath())
	// screencastForm, err := render.FromHTML(
	// 	"screencasts/new.html",
	// 	r.r,
	// 	map[string]interface{}{
	// 		"screencast": &models.Screencast{},
	// 	})
	// if err != nil {
	// 	return err
	// }
	// view, err := components.Base(authenticityToken, r.manifest, screencastForm)
	// if err != nil {
	// 	return err
	// }
	screencast := &models.Screencast{}
	c.Set("screencast", screencast)
	c.Set("nav", navToHTML(c))
	return c.Render(200, r.r.HTML("screencasts/new.html"))
	// TODO: get forms working! Need to figure out the authenticity_token
	// return c.Render(200, render.EltToRenderer(view))
}

// Create adds a Screencast to the DB. This function is mapped to the
// path POST /screencasts
func (r Resource) Create(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	// Bind screencast to the html form elements
	if err := c.Bind(screencast); err != nil {
		return err
	}

	videoFile, err := c.File("Video")
	if err != nil {
		return err
	}
	defer videoFile.Close()

	ytSvc, err := r.ytConfig.Service(c)
	if err != nil {
		return err
	}
	// ytResponseID, ytUploadErr := youtube.UploadToYoutube(
	_, ytUploadErr := youtube.UploadToYoutube(
		ytSvc,
		videoFile,
		r.ytConfig.ChannelID(),
		screencast.Title,
	)
	if ytUploadErr != nil {
		return ytUploadErr
	}
	// TODO: save ytResponseID somewhere!

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(screencast)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Error(500, fmt.Errorf("Errors creating the screencast! %s", verrs))
	}

	return r.Redirect(c)
}

// Edit renders a edit form for a Screencast. This function is
// mapped to the path GET /screencasts/{_id}/edit
func (r Resource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	if err := tx.Find(screencast, c.Param("_id")); err != nil {
		return c.Error(404, err)
	}

	c.Set("nav", navToHTML(c))
	c.Set("screencast", screencast)
	return c.Render(200, r.r.HTML("screencasts/edit.html"))
}

// Update changes a Screencast in the DB
func (r Resource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	if err := tx.Find(screencast, c.Param("_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Screencast to the html form elements
	if err := c.Bind(screencast); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(screencast)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Error(500, fmt.Errorf("Error updating! %s", verrs))
	}

	return r.Redirect(c)
}

// Destroy deletes a Screencast from the DB. This function is mapped
// to the path DELETE /screencasts/{_id}
func (r Resource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	// To find the Screencast the parameter _id is used.
	if err := tx.Find(screencast, c.Param("_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(screencast); err != nil {
		return c.Error(500, err)
	}

	return r.Redirect(c)
}
