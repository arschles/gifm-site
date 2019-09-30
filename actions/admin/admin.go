package admin

import (
	"fmt"

	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/views/admin"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// HomeRoute creates the route necessary to serve GET /admin
func HomeRoute(manifest *assets.Manifest) buffalo.Handler {
	return func(c buffalo.Context) error {
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
		if err := q.Order("episode_num desc").All(screencasts); err != nil {
			return err
		}

		view, err := admin.Home(c, manifest, screencasts)
		if err != nil {
			return err
		}
		return c.Render(200, render.EltToRenderer(view))
	}
}
