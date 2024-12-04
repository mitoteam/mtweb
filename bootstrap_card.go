package mtweb

import "github.com/mitoteam/dhtml"

//Some Bootstrap Css Framework helpers
//https://getbootstrap.com/docs/5.3/getting-started/introduction/

type (
	Card struct {
		header *dhtml.HtmlPiece
		body   *dhtml.HtmlPiece
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*Card)(nil)

func NewCard() *Card {
	return &Card{}
}

// Appends something to header
func (c *Card) Header(v any) *Card {
	c.GetHeader().Append(v)
	return c
}

func (c *Card) GetHeader() *dhtml.HtmlPiece {
	if c.header == nil {
		c.header = dhtml.NewHtmlPiece() //empty piece
	}

	return c.header
}

// Appends something to body
func (c *Card) Body(v any) *Card {
	c.GetBody().Append(v)
	return c
}

// Pointer to card's body
func (c *Card) GetBody() *dhtml.HtmlPiece {
	if c.body == nil {
		c.body = dhtml.NewHtmlPiece() //empty piece
	}

	return c.body
}

// region Rendering
func (c *Card) GetTags() dhtml.TagsList {
	root := dhtml.Div().Class("card")

	if c.header != nil {
		root.Append(dhtml.Div().Class("card-header").Append(c.header))
	}

	if c.body != nil {
		root.Append(
			dhtml.Div().Class("card-body").Append(c.body),
		)
	}

	return root.GetTags()
}

//endregion
