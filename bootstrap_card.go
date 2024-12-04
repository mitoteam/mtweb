package mtweb

import "github.com/mitoteam/dhtml"

//https://getbootstrap.com/docs/5.3/components/card/

type (
	Card struct {
		header dhtml.HtmlPiece
		body   dhtml.HtmlPiece
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*Card)(nil)

func NewCard() *Card {
	return &Card{}
}

// Appends something to header
func (c *Card) Header(v any) *Card {
	c.header.Append(v)
	return c
}

func (c *Card) GetHeader() *dhtml.HtmlPiece {
	return &c.header
}

// Appends something to body
func (c *Card) Body(v any) *Card {
	c.body.Append(v)
	return c
}

// Pointer to card's body
func (c *Card) GetBody() *dhtml.HtmlPiece {
	return &c.body
}

// region Rendering
func (c *Card) GetTags() dhtml.TagsList {
	root := dhtml.Div().Class("card")

	if !c.header.IsEmpty() {
		root.Append(dhtml.Div().Class("card-header").Append(c.header))
	}

	if !c.body.IsEmpty() {
		root.Append(
			dhtml.Div().Class("card-body").Append(c.body),
		)
	}

	return root.GetTags()
}

//endregion
