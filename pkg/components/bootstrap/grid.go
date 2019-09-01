package bootstrap

import (
	"io"

	"github.com/arschles/go-in-5-minutes-site/pkg/render"
)

type Row struct {
	opts render.TagOpts
	cols []Col
}

func (r Row) ToHTML() (io.Reader, error) {
	opts := render.TagOpts{
		"class": "row",
	}

	colElts := make([]render.Elt, len(r.cols))
	for i, col := range r.cols {
		colElts[i] = render.Tag("div", col.opts, col.elts)
	}
	return render.Tag("div", render.MergeTagOpts(r.opts, opts, "class"), colElts...).ToHTML()
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

func (c Col) WithChild(e render.Elt) Col {
	c.elts = append(c.elts, e)
	return c
}

type Grid struct {
	opts render.TagOpts
	rows []Row
}

func NewGrid() Grid {
	return Grid{rows: nil}
}

// func (g Grid) WithRow(r Row) Grid

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