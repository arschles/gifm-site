package bootstrap

import (
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

// FormGroup returns a TagBuilder for building up a "class":"form-group"
// div
func FormGroup(opts render.TagOpts) render.TagBuilder {
	return render.NewTag("div").WithOpt("class", "form-group")
}

// Label creates a new <label> tag
func Label(forID, text string, opts render.TagOpts) render.TagBuilder {
	return tags.Label(render.TagOpts{"for": "ep-num"}).WithChild(
		render.Text(text),
	)
}

// Input creates a new <input> tag
func Input(inputType, id, placeholder string, opts render.TagOpts) render.TagBuilder {
	mergedOpts := render.MergeTagOpts(opts, render.TagOpts{
		"type":        inputType,
		"class":       "form-control",
		"id":          id,
		"placeholder": placeholder,
	}, "class")
	return render.NewTag("input").WithOpts(mergedOpts)
}

// TextArea creates a new <textarea> tag
func TextArea(opts render.TagOpts) render.TagBuilder {
	mergedOpts := render.MergeTagOpts(opts, render.TagOpts{
		"class":    "form-control",
		"required": "",
	}, "class")
	return render.NewTag("textarea").WithOpts(mergedOpts)
}
