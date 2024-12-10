package mtweb

import "github.com/mitoteam/dhtml"

type (
	CardListElement struct {
		cards []*CardElement
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*CardListElement)(nil)

func NewCardList() *CardListElement {
	return &CardListElement{}
}

func (e *CardListElement) Add(card *CardElement) *CardListElement {
	e.cards = append(e.cards, card)
	return e
}

func (e *CardListElement) GetTags() dhtml.TagsList {
	root := dhtml.Div().Class("card-list row row-cols-md-2 g-3")

	for _, card := range e.cards {
		root.Append(dhtml.Div().Class("col").Append(card))
	}

	return root.GetTags()
}
