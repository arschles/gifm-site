package actions

import (
	"os"

	"github.com/arschles/gifm-site/pkg/security"
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
	security.PopulateSession(c, user)
	// Do something with the user, maybe register them/sign them in
	return c.Redirect(302, "/admin")
}

func authorizeMiddleware() func(buffalo.Handler) buffalo.Handler {
	return func(next buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			if !security.IsAdmin(c) {
				return c.Redirect(302, "/auth/github")
			}
			return next(c)
		}
	}
}
