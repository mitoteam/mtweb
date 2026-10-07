package mtweb

import (
	"time"

	"github.com/mitoteam/dhtml"
)

// Clock icon with datetime
type TimestampElement struct {
	format  string
	icon    string
	ts      time.Time
	classes dhtml.Classes
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*TimestampElement)(nil)

const IconTimestamp = "clock"

func NewTimestamp(ts time.Time) *TimestampElement {
	return &TimestampElement{
		ts:     ts,
		icon:   IconTimestamp,
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
