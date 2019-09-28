package grifts

import (
	"github.com/arschles/gifm-site/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
