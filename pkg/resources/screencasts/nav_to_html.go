package screencasts

import (
	"html/template"
	"io/ioutil"

	"github.com/arschles/gifm-site/views/components"
	"github.com/gobuffalo/buffalo"
)

func navToHTML(c buffalo.Context) func() (template.HTML, error) {
	return func() (template.HTML, error) {
		navHTMLReader, err := components.Nav(c).ToHTML()
		if err != nil {
			return "", err
		}
		navHTMLBytes, err := ioutil.ReadAll(navHTMLReader)
		if err != nil {
			return "", err
		}
		return template.HTML(string(navHTMLBytes)), nil
	}
}
