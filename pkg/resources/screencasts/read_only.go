package screencasts

import (
	"fmt"

	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/views/components"
	"github.com/arschles/gifm-site/views/components/bootstrap"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/views"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// ReadOnlyResource has the routes in it for read-only
type ReadOnlyResource struct {
	manifest *assets.Manifest
}

// NewReadOnlyResource creates a new ReadOnly instance
func NewReadOnlyResource(manifest *assets.Manifest) ReadOnlyResource {
	return ReadOnlyResource{manifest: manifest}
}

// List gets all Screencasts
// This function is mapped to the path:
// GET /screencasts
func (r ReadOnlyResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	screencasts := &models.Screencasts{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Screencasts from the DB
	if err := q.All(screencasts); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	view, err := views.Screencasts(c, r.manifest, screencasts)
	if err != nil {
		return err
	}

	return c.Render(200, render.EltToRenderer(view))
}

// Show gets the data for one Screencast. This function is mapped to
// the path GET /screencasts/{screencast_id}
func (r ReadOnlyResource) Show(c buffalo.Context) error {
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

	view, err := components.Base(
		c,
		r.manifest,
		bootstrap.NewGrid(render.EmptyOpts()).WithRow(
			bootstrap.NewRow(render.EmptyOpts()).WithCol(
				bootstrap.NewCol(render.EmptyOpts()).WithChild(
					render.Text(screencast.Title),
				),
			),
		),
	)
	if err != nil {
		return c.Error(500, err)
	}
	return c.Render(200, render.EltToRenderer(view))
}
