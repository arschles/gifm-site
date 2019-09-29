package screencasts

import (
	"io/ioutil"
	"html/template"

	"github.com/arschles/gifm-site/views/components"
)

func navToHTML() (template.HTML, error) {
	navHTMLReader, err := components.Nav().ToHTML()
	if err != nil {
		return "", err
	}
	navHTMLBytes, err := ioutil.ReadAll(navHTMLReader)
	if err != nil {
		return "", err
	}
	return template.HTML(string(navHTMLBytes)), nil
}
