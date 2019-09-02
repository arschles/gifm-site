package actions

import (
	"fmt"

	"github.com/arschles/go-in-5-minutes-site/models"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/views"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
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
	c.Set("screencasts", screencasts)
	homeView, err := views.Home(parsedManifest)
	if err != nil {
		return err
	}
	return c.Render(200, render.EltToRenderer(homeView))
}
