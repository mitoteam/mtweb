package mtweb

import "github.com/mitoteam/dhtml"

func RenderEmptyValue(label string) *dhtml.Tag {
	return dhtml.Span().Class("text-muted").Append("[" + label + "]")
}

// =================== JustifiedLR =========================

// couple of <div> tags justified by applying .d-flex and .justify-content-between classes
type JustifiedLRElement struct {
	l dhtml.HtmlPiece
	r dhtml.HtmlPiece
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*JustifiedLRElement)(nil)

func NewJustifiedLR() *JustifiedLRElement {
	return &JustifiedLRElement{}
}

func (e *JustifiedLRElement) GetL() *dhtml.HtmlPiece {
	return &e.l
}

func (e *JustifiedLRElement) L(v any) *JustifiedLRElement {
	e.l.Append(v)
	return e
}

func (e *JustifiedLRElement) GetR() *dhtml.HtmlPiece {
	return &e.r
}

func (e *JustifiedLRElement) R(v any) *JustifiedLRElement {
	e.r.Append(v)
	return e
}

func (j *JustifiedLRElement) GetTags() dhtml.TagsList {
	return dhtml.Div().Class([]string{"d-flex", "justify-content-between"}).
		Append(dhtml.Div().Append(j.l)).
		Append(dhtml.Div().Append(j.r)).
		GetTags()
}
