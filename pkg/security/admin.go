package security

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/goth"
)

func LoggedInUser(c buffalo.Context) string {
	iface := c.Session().Get("current_user")
	if iface == nil {
		return ""
	}
	str, ok := iface.(string)
	if !ok {
		return ""
	}
	return str
}

func PopulateSession(c buffalo.Context, user goth.User) {
	c.Session().Set("current_user", user.NickName)
}

func IsAdmin(c buffalo.Context) bool {
	return LoggedInUser(c) == "arschles"
}
