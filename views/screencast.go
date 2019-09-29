package views

import (
	"bytes"

	"github.com/arschles/gifm-site/models"
	"github.com/arschles/gifm-site/pkg/render"
	"github.com/arschles/gifm-site/views/components/bootstrap"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Screencast renders a single screencast
func Screencast(screencast models.Screencast) (render.Elt, error) {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(screencast.Markdown), &buf); err != nil {
		return nil, err
	}
	mdStr := string(buf.Bytes())

	stdRow := bootstrap.NewRow(render.TagOpts{"class": "mt-4"})
	stdCol := bootstrap.NewCol(render.TagOpts{"class": "col-sm-10"})

	return bootstrap.NewGrid(render.TagOpts{"class": "mt-4"}).WithRow(
		stdRow.WithCol(
			stdCol.WithChild(
				render.NewTag("h1").WithChild(
					render.Text(screencast.Title),
				),
			),
		),
	).WithRow(
		stdRow.WithCol(
			stdCol.WithChild(
				render.NewTag("div").
					WithOpt("class", "shadow-none p-3 mb-5 bg-light rounded").
					WithChild(render.Text(mdStr)),
			),
		),
	), nil
	// return tags.Div(render.TagOpts{"class": "row mt-5"},
	// 	tags.Div(render.TagOpts{"class": "col-sm-12"},
	// 		bootstrap.Card(screencast.Title, screencast.Intro),
	// 		bootstrap.Card("Markdown", string(buf.Bytes())),
	// 	),
	// ), nil
}
