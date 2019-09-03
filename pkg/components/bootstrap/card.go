package bootstrap

import (
	"io"

	"github.com/arschles/go-in-5-minutes-site/pkg/render"
)

type CardBuilder struct {
	render.Elt
	title string
	text  string
}

func (c CardBuilder) ToHTML() (io.Reader, error) {
	allOpts := render.TagOpts{"class": "card", "style": "width:10rem;"}

	titleElt := render.NewTag("h5").WithOpt("class", "card-title").WithText(c.title)
	textElt := render.NewTag("p").WithOpt("class", "card-text").WithText(c.text)

	cardGroup := render.NewTag("div").WithOpt("class", "card-group")
	return cardGroup.WithChild(
		render.NewTag("div").WithOpts(allOpts).WithChild(
			render.NewTag("div").WithOpt("class", "card-body").WithChildren(titleElt, textElt),
		),
	).ToHTML()
}

// Card returns a render.Elt that holds a bootstrap card.
//
// See: https://getbootstrap.com/docs/4.3/components/card/
func Card(title string, text string, otherElts ...render.Elt) *CardBuilder {
	return &CardBuilder{
		title: title,
		text:  text,
	}
}
