package helpers

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/tags"
)

func addOpt(opts tags.Options, k string, val interface{}) tags.Options {
	if v, ok := opts[k]; ok {
		newVal := fmt.Sprintf("%s %s", v, val)
		opts[k] = newVal
	}
	return opts
}

func optsAsString(opts tags.Options) string {
	attrs := []string{}
	for k, v := range opts {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, k, v))
	}
	return strings.Join(attrs, " ")
}

func Tag(name string, opts tags.Options, h template.HTML) template.HTML {
	return template.HTML(fmt.Sprintf(`<%s %s>%s</%s>`, name, optsAsString(opts), h))
}

// TODO: make these like the formFor and form helpers.
// You can put the body of the div (and ther other thingies) inside
// the function call that way
func Div(opts tags.Options, h template.HTML) template.HTML {
	return Tag("div", opts, h)
}

func Container(opts tags.Options, h template.HTML) template.HTML {
	opts["class"] = "container"
	return Div(addOpt(opts, "class", "container"), h)
}

func Row(opts tags.Options, h template.HTML) template.HTML {
	return Div(addOpt(opts, "class", "row"), h)
}

func Text(s string) template.HTML {
	return template.HTML(s)
}
