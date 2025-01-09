package mtweb

import (
	"time"

	"github.com/mitoteam/dhtml"
)

// Count label
func RenderCount(count int64, title string) (out dhtml.HtmlPiece) {
	out.Append(
		Icon("abacus").Label(count).ElementClass("p-1 border border-dark-subtle fw-bold bg-info-subtle").Title(title),
	)

	return out
}

// Bootstrap styled table
func NewTable() *dhtml.TableElement {
	return dhtml.NewTable().Class("table table-hover table-sm").EmptyLabel("nothing here yet")
}

// Table count label
func RenderTableCount(table *dhtml.TableElement, title string) (out dhtml.HtmlPiece) {
	cnt := int64(table.RowCount())

	if cnt > 0 {
		out.Append(
			dhtml.Div().Class("text-end").Append(RenderCount(cnt, title)),
		)
	}

	return out
}

// Clock icon with datetime
func RenderTimestamp(ts time.Time) (out dhtml.HtmlPiece) {
	out.Append(
		Icon("clock").Label(ts.Format(time.DateTime)).ElementClass("small text-muted"),
	)

	return out
}
