package mtweb

import "github.com/mitoteam/dhtml"

//Some Bootstrap Css Framework helpers
//https://getbootstrap.com/docs/5.3/getting-started/introduction/

type (
	Card struct {
		Header *dhtml.ElementList
		body   *dhtml.ElementList
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*Card)(nil)

func (c *Card) AppendList(list *dhtml.ElementList) *Card {
	if c.body == nil {
		c.body = dhtml.Piece(list)
	} else {
		c.body.AppendList(list)
	}

	return c
}

func (c *Card) Append(element dhtml.ElementI) *Card {
	return c.AppendList(dhtml.NewElementList().Append(element))
}

// region Rendering
func (c *Card) GetTags() dhtml.TagsList {
	root := dhtml.Div().Class("card")

	if c.Header != nil {
		root.Append(dhtml.Div().Class("card-header").Append(c.Header))
	}

	if c.body != nil {
		root.Append(
			dhtml.Div().Class("card-body").AppendList(c.body),
		)
	}

	return dhtml.TagsList{root}
}

//endregion
