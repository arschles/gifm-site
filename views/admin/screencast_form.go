package admin

import (
	bts "github.com/arschles/go-in-5-minutes-site/pkg/components/bootstrap"
	"github.com/arschles/go-in-5-minutes-site/pkg/forms"
	"github.com/arschles/go-in-5-minutes-site/pkg/render"
	"github.com/arschles/go-in-5-minutes-site/pkg/tags"
)

// ScreencastForm returns a TagBuilder that contains the entire for for creating
// a new screencast
//
// TODO: make this optionally take a screencast so you can edit the screencast
// too
func ScreencastForm(authenticityToken, formAction string) render.TagBuilder {
	return render.NewTag("form").WithOpts(render.TagOpts{
		"method": "POST",
		"action": formAction,
	}).WithChildren(
		tags.Input(render.TagOpts{
			"name":  forms.TokenFieldName,
			"value": authenticityToken,
			"type":  "hidden",
		}),
		bts.FormGroup(render.EmptyOpts()).WithChild(
			bts.Label("episode_num", "Episode Number", render.EmptyOpts()),
		).WithChild(
			bts.Input("text", "episode_num", "Episode Number", render.EmptyOpts()),
		),
		bts.FormGroup(render.EmptyOpts()).WithChild(
			bts.Label("title", "Title", render.EmptyOpts()),
		).WithChild(
			bts.Input("text", "title", "Title", render.TagOpts{
				"class": "form-control form-control-lg",
			}),
		),
		bts.FormGroup(render.EmptyOpts()).WithChild(
			bts.Label("intro", "Intro", render.EmptyOpts()),
		).WithChild(
			bts.TextArea(render.TagOpts{
				"id":       "intro",
				"rows":     "4",
				"required": "",
			}),
		),
		bts.FormGroup(render.EmptyOpts()).WithChild(
			bts.Label("markdown", "Markdown Description", render.EmptyOpts()),
		).WithChild(
			bts.TextArea(render.TagOpts{
				"id":       "Markdown",
				"rows":     "8",
				"required": "",
			}),
		),
		render.NewTag("button").WithOpts(render.TagOpts{
			"type":  "submit",
			"class": "btn btn-primary",
		}).WithChild(
			render.Text("Create"),
		),
	)
}
