package tags

import "github.com/arschles/gifm-site/pkg/render"

// Label creates a new <label> tag
func Label(opts render.TagOpts) render.TagBuilder {
	return render.NewTag("label").WithOpts(opts)
}

// Input creates a new <input> tag
func Input(opts render.TagOpts) render.TagBuilder {
	return render.NewTag("input").WithOpts(opts)
}
