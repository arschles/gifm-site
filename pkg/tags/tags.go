package tags

import (
	"github.com/arschles/gifm-site/pkg/assets"
	"github.com/arschles/gifm-site/pkg/render"
)

func Link(opts render.TagOpts) render.Elt {
	return Empty("link", opts)
}

func Meta(opts render.TagOpts) render.Elt {
	return Empty("meta", opts)
}

func A(href string, opts render.TagOpts, text string) render.Elt {
	opts = render.MergeTagOpts(opts, render.TagOpts{"href": href})
	return render.Tag("a", opts, render.Text(text))
}

func Div(opts render.TagOpts, contents ...render.Elt) render.Elt {
	return render.Tag("div", opts, contents...)
}

func Span(opts render.TagOpts, contents ...render.Elt) render.Elt {
	return render.Tag("span", opts, contents...)
}

func P(opts render.TagOpts, contents ...render.Elt) render.Elt {
	return render.Tag("p", opts, contents...)
}

func Small(opts render.TagOpts, contents ...render.Elt) render.Elt {
	return render.Tag("small", opts, contents...)
}

func JS(manifest *assets.Manifest, name string) (render.Elt, error) {
	fname, err := manifest.FullyQualified(name)
	if err != nil {
		return nil, err
	}
	return render.Tag("script", render.TagOpts{
		"type": "text/javascript",
		"src":  fname,
	}), nil
}

func CSS(manifest *assets.Manifest, name string) (render.Elt, error) {
	fname, err := manifest.FullyQualified(name)
	if err != nil {
		return nil, err
	}
	return Link(render.TagOpts{
		"rel":  "stylesheet",
		"href": fname,
	}), nil
}
