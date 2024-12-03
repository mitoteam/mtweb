package mtweb

import "github.com/mitoteam/dhtml"

//Some Bootstrap Css Framework helpers
//https://getbootstrap.com/docs/5.3/getting-started/introduction/

type (
	Card struct {
		Title string

		body dhtml.ElementsList
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*Card)(nil)

func (c *Card) GetTags() dhtml.TagsList {
	root := dhtml.Div().Class("card")

	if c.Title != "" {
		root.Append(dhtml.Div().Class("card-header").Text(c.Title))
	}

	if c.body != nil {
		root.Append(
			dhtml.Div().Class("card-body").AppendList(c.body),
		)
	}

	return dhtml.TagsList{root}
}

func (c *Card) AppendList(list dhtml.ElementsList) *Card {
	if c.body == nil {
		c.body = make(dhtml.ElementsList, 0)
	}

	c.body = append(c.body, list...)

	return c
}

func (c *Card) Append(element dhtml.ElementI) *Card {
	return c.AppendList(dhtml.ElementsList{element})
}
