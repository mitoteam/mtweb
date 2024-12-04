package mtweb

import "github.com/mitoteam/dhtml"

//Some Bootstrap Css Framework helpers
//https://getbootstrap.com/docs/5.3/getting-started/introduction/

type (
	Card struct {
		Header *dhtml.HtmlPiece
		body   *dhtml.HtmlPiece
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*Card)(nil)

// Appends something to body
func (c *Card) Append(v any) *Card {
	c.Body().Append(v)
	return c
}

// Pointer to elements body
func (c *Card) Body() *dhtml.HtmlPiece {
	if c.body == nil {
		c.body = dhtml.NewHtmlPiece() //empty piece
	}

	return c.body
}

// region Rendering
func (c *Card) GetTags() dhtml.TagsList {
	root := dhtml.Div().Class("card")

	if c.Header != nil {
		root.Append(dhtml.Div().Class("card-header").Append(c.Header))
	}

	if c.body != nil {
		root.Append(
			dhtml.Div().Class("card-body").Append(c.body),
		)
	}

	return dhtml.TagsList{root}
}

//endregion
