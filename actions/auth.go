package actions

import (
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func setupGoth(app *buffalo.App) {
	// gothic.Store = app.SessionStore
	ghKey := os.Getenv("GITHUB_KEY")
	ghSecret := os.Getenv("GITHUB_SECRET")
	callback := os.Getenv("AUTH_CALLBACK")
	goth.UseProviders(
		github.New(ghKey, ghSecret, callback),
	)
}

func AuthCallback(c buffalo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	c.Session().Set("current_user_id", user.UserID)
	// Do something with the user, maybe register them/sign them in
	return c.Redirect(302, "/admin")
}

func authorizeMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			return c.Redirect(302, "/auth/github")
		}
		return next(c)
	}
}
