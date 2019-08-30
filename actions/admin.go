package actions

import "github.com/gobuffalo/buffalo"

func adminHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("admin/index.plush.html"))
}