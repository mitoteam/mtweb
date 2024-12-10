package mtweb

import "github.com/mitoteam/dhtml"

// couple of <div> tags justified by applying .d-flex and .justify-content-between classes
type JustifiedLR struct {
	l dhtml.HtmlPiece
	r dhtml.HtmlPiece
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*JustifiedLR)(nil)

func NewJustifiedLR() *JustifiedLR {
	return &JustifiedLR{}
}

func (j *JustifiedLR) GetL() *dhtml.HtmlPiece {
	return &j.l
}

func (j *JustifiedLR) L(v any) *JustifiedLR {
	j.l.Append(v)
	return j
}

func (j *JustifiedLR) GetR() *dhtml.HtmlPiece {
	return &j.r
}

func (j *JustifiedLR) R(v any) *JustifiedLR {
	j.r.Append(v)
	return j
}

func (j *JustifiedLR) GetTags() dhtml.TagsList {
	return dhtml.Div().Class([]string{"d-flex", "justify-content-between"}).
		Append(dhtml.Div().Append(j.l)).
		Append(dhtml.Div().Append(j.r)).
		GetTags()
}
