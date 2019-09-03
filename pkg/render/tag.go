package render

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func Tag(name string, opts TagOpts, contents ...Elt) Elt {
	return tag{name: name, opts: opts, children: contents}
}

func NewTag(name string) TagBuilder {
	return TagBuilder{name: name, opts: EmptyOpts(), children: nil}
}

type tag struct {
	name     string
	opts     TagOpts
	children []Elt
}

func (e tag) ToHTML() (io.Reader, error) {
	children := make([]string, len(e.children))
	for i, child := range e.children {
		childReader, err := child.ToHTML()
		if err != nil {
			return nil, err
		}
		childBytes, err := ioutil.ReadAll(childReader)
		if err != nil {
			return nil, err
		}
		children[i] = string(childBytes)
	}
	optsStr := ""
	if len(e.opts) > 0 {
		optsStr = " " + e.opts.String()
	}
	toWrite := fmt.Sprintf(`<%s%s>
%s
</%s>`,
		e.name,
		optsStr,
		strings.Join(children, "\n"),
		e.name,
	)
	return bytes.NewBuffer([]byte(toWrite)), nil
}

type TagBuilder struct {
	name     string
	opts     TagOpts
	children []Elt
}

func (t TagBuilder) WithOpt(name string, val interface{}) TagBuilder {
	t.opts[name] = val
	return t
}

func (t TagBuilder) WithOpts(opts TagOpts) TagBuilder {
	t.opts = opts
	return t
}

func (t TagBuilder) WithChild(e Elt) TagBuilder {
	t.children = append(t.children, e)
	return t
}

func (t TagBuilder) WithChildren(elts ...Elt) TagBuilder {
	ret := t
	for _, elt := range elts {
		ret = ret.WithChild(elt)
	}
	return ret
}

func (t TagBuilder) WithText(s string) TagBuilder {
	return t.WithChild(Text(s))
}

func (t TagBuilder) ToHTML() (io.Reader, error) {
	return Tag(t.name, t.opts, t.children...).ToHTML()
}
