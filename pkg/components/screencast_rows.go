package components

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/arschles/go-in-5-minutes-site/models"
)

func ScreencastRows(screencastsPtr *models.Screencasts) func() template.HTML {
	return func() template.HTML {
		screencasts := *screencastsPtr
		elts := []string{`<div class="row">Screencasts</div>`}
		if len(screencasts) < 1 {
			elts = append(elts, `<div class="row">No Screencasts!</div>`)
		} else {
			for _, screencast := range screencasts {
				elts = append(elts, fmt.Sprintf(`<div class="row">%s</div>`, screencast.Title))
			}
		}
		return template.HTML(strings.Join(elts, "\n"))
	}
}
