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

// ================== timestamp element ====================
// Clock icon with datetime
type TimestampElement struct {
	format  string
	icon    string
	ts      time.Time
	classes dhtml.Classes
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*TimestampElement)(nil)

func NewTimestamp(ts time.Time) *TimestampElement {
	return &TimestampElement{
		ts:     ts,
		icon:   "clock",
		format: time.DateTime,
	}
}

// Element classes.
func (e *TimestampElement) Class(v ...any) *TimestampElement {
	e.classes.Add(v...)
	return e
}

func (e *TimestampElement) SmallMuted() *TimestampElement {
	e.classes.Add("small", "text-muted")
	return e
}

// Set date and time format (default is time.DateTime).
func (e *TimestampElement) Format(tsFormat string) *TimestampElement {
	e.format = tsFormat
	return e
}

// icon name. empty = no icon. default "clock".
func (e *TimestampElement) Icon(icon string) *TimestampElement {
	e.icon = icon
	return e
}

func (e *TimestampElement) GetTags() dhtml.TagList {
	rootTag := dhtml.Div().Class(e.classes)

	if e.icon != "" {
		rootTag.Append(Icon(e.icon).Label(e.ts.Format(e.format)))
	} else {
		rootTag.Append(e.ts.Format(e.format))
	}

	return rootTag.GetTags()
}
