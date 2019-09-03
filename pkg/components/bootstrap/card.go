package bootstrap

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
)

// Card returns a render.Elt that holds a bootstrap card.
//
// See: https://getbootstrap.com/docs/4.3/components/card/
func Card(title string, text string, opts ...render.TagOpts) render.Elt {
	titleElt := render.NewTag("h5").WithOpt("class", "card-title").WithText(title)
	textElt := render.NewTag("p").WithOpt("class", "card-text").WithText(text)
	allOpts := render.TagOpts{"class": "card", "style": "width:10rem;"}
	for _, opt := range opts {
		allOpts = render.MergeTagOpts(opt, allOpts, "class", "style")
	}
	return render.NewTag("div").WithOpt("class", "card-group").WithChild(
		render.NewTag("div").WithOpts(allOpts).WithChild(
			render.NewTag("div").WithOpt("class", "card-body").WithChildren(titleElt, textElt),
		),
	)
}
