package resources

import (
	"github.com/gobuffalo/buffalo"
)

// RedirKey is the key used in query strings to indicate where to redirect
const RedirKey = "redir"

// Base is the base of all resources. Compose it into your structs
type Base struct {
	basePath string
}

// NewBase creates a new base with the given base path
func NewBase(basePath string) Base {
	return Base{basePath: basePath}
}

// Redirect uses c to send a 302 Temporary Redirect back to the client.
//
// It uses RedirKey to figure out where to redirect. If there is no RedirKey
// present, it instead uses the base path it's configured with
func (b Base) Redirect(c buffalo.Context) error {
	return c.Redirect(302, b.StringParam(c, RedirKey, b.basePath))
}

// StringParam returns the value of the parameter "name" from c,
// or defalt if it doesn't exist
func (b Base) StringParam(c buffalo.Context, name, defalt string) string {
	ret := c.Param(name)
	if ret == "" {
		return defalt
	}
	return ret
}
