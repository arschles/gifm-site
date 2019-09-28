package bootstrap

import (
	"io"

	"github.com/arschles/gifm-site/pkg/render"
)

type Row struct {
	opts render.TagOpts
	cols []Col
}

func (r Row) ToHTML() (io.Reader, error) {
	opts := render.MergeTagOpts(r.opts, render.TagOpts{
		"class": "row",
	}, "class")

	colElts := make([]render.Elt, len(r.cols))
	for i, col := range r.cols {
		colElts[i] = render.Tag("div", col.opts, col.elts...)
	}
	return render.Tag("div", opts, colElts...).ToHTML()
}

func (r Row) WithCol(c Col) Row {
	r.cols = append(r.cols, c)
	return r
}

func NewRow(opts render.TagOpts, cols ...Col) Row {
	return Row{opts: opts, cols: cols}
}

type Col struct {
	opts render.TagOpts
	elts []render.Elt
}

func NewCol(opts render.TagOpts) Col {
	return Col{opts: opts, elts: nil}
}

func (c Col) WithChild(e render.Elt) Col {
	c.elts = append(c.elts, e)
	return c
}

type Grid struct {
	opts render.TagOpts
	rows []Row
}

func NewGrid(opts render.TagOpts) Grid {
	return Grid{opts: opts, rows: nil}
}

func (g Grid) WithRow(r Row) Grid {
	g.rows = append(g.rows, r)
	return g
}

func (g Grid) ToHTML() (io.Reader, error) {
	rowElts := make([]render.Elt, len(g.rows))
	for i, row := range g.rows {
		rowElts[i] = row
	}
	finalElt := render.Tag(
		"div",
		render.MergeTagOpts(g.opts, render.TagOpts{"class": "container"}, "class"),
		rowElts...,
	)
	return finalElt.ToHTML()
}
