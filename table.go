package mtweb

import (
	"github.com/mitoteam/dhtml"
)

// Bootstrap styled table with some some utilities
type TableElement struct {
	dhtml.TableElement

	countTitle string // element count is rendered if this is not empty
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*TableElement)(nil)

func NewTable() *TableElement {
	t := &TableElement{
		TableElement: *dhtml.NewTable(),
	}

	t.Class("table table-hover table-sm").
		BodyClass("table-group-divider").
		EmptyLabel("nothing here yet")

	return t
}

// Sets title for count element. Count element is not rendered if title is empty string.
func (t *TableElement) CountTitle(v string) *TableElement {
	t.countTitle = v
	return t
}

func (t *TableElement) GetTags() (out dhtml.TagList) {
	p := dhtml.NewHtmlPiece()

	if t.countTitle != "" {
		if cnt := int64(t.RowCount()); cnt > 0 {
			p.Append(
				dhtml.Div().
					Class("text-end").
					Append(
						Icon(FaIconCount).Label(cnt).ElementClass("p-1 border border-dark-subtle fw-bold bg-info-subtle").Title(t.countTitle),
					),
			)
		}
	}

	p.Append(t.TableElement.GetTags())

	return p.GetTags()
}
