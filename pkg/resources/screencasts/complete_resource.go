package screencasts

import (
	"fmt"

	"github.com/arschles/go-in-5-minutes-site/models"
	"github.com/arschles/go-in-5-minutes-site/pkg/assets"
	"github.com/arschles/go-in-5-minutes-site/pkg/components"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/resources"
	"github.com/arschles/go-in-5-minutes-site/views/admin"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// Resource is a complete buffalo.Resource for screencasts
type Resource struct {
	ReadOnlyResource
	resources.Base
}

// NewResource creates a new Resource for screencasts
func NewResource(basePath string, manifest *assets.Manifest) Resource {
	return Resource{
		ReadOnlyResource: NewReadOnlyResource(manifest),
		Base:             resources.NewBase(basePath),
	}
}

// New renders the form for creating a new Screencast.
// This function is mapped to the path GET /screencasts/new
func (r Resource) New(c buffalo.Context) error {
	view, err := components.Base(r.manifest, admin.ScreencastForm())
	if err != nil {
		return err
	}
	return c.Render(200, render.EltToRenderer(view))
}

// Create adds a Screencast to the DB. This function is mapped to the
// path POST /screencasts
func (r Resource) Create(c buffalo.Context) error {
	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	// Bind screencast to the html form elements
	if err := c.Bind(screencast); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

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
// mapped to the path GET /screencasts/{screencast_id}/edit
func (r Resource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
		return c.Error(404, err)
	}

	return r.Redirect(c)
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

	if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
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
// to the path DELETE /screencasts/{screencast_id}
func (r Resource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Screencast
	screencast := &models.Screencast{}

	// To find the Screencast the parameter screencast_id is used.
	if err := tx.Find(screencast, c.Param("screencast_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(screencast); err != nil {
		return c.Error(500, err)
	}

	return r.Redirect(c)
}
