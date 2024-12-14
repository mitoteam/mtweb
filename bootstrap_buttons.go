package mtweb

import (
	"html"

	"github.com/mitoteam/dhtml"
)

//https://getbootstrap.com/docs/5.3/components/buttons/

var (
	BtnSizeClasses = []string{
		"btn-sm", "btn-md", "btn-lg",
	}

	BtnVariantClasses = []string{
		"btn-outline-primary", "btn-outline-secondary", "btn-outline-success", "btn-outline-danger", "btn-outline-warning",
		"btn-outline-info", "btn-outline-light", "btn-outline-dark",
		"btn-primary", "btn-secondary", "btn-success", "btn-danger", "btn-warning", "btn-info", "btn-light", "btn-dark",
	}
	DefaultBtnVariantClass = BtnVariantClasses[0]
)

type BtnElement struct {
	tag *dhtml.Tag
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*BtnElement)(nil)

func NewBtn() *BtnElement {
	return &BtnElement{
		tag: dhtml.NewTag("a"),
	}
}

func (e *BtnElement) Href(url string) *BtnElement {
	e.tag.Attribute("href", url)
	return e
}

func (e *BtnElement) Label(v any) *BtnElement {
	e.tag.Append(v)
	return e
}

func (e *BtnElement) Title(title string) *BtnElement {
	e.tag.Title(title)
	return e
}

func (e *BtnElement) Class(v any) *BtnElement {
	e.tag.Class(v)
	return e
}

func (e *BtnElement) GetTags() dhtml.TagsList {
	e.tag.GetClasses().
		Add("btn text-nowrap").
		AddFromSet(DefaultBtnVariantClass, BtnVariantClasses)

	return e.tag.GetTags()
}

// ======== Btn helpers ==========

func NewEditBtn(href string) *BtnElement {
	return NewBtn().Href(href).Title("Edit").Class("btn-sm px-1").
		Label(Icon("edit"))
}

func NewDeleteBtn(href, confirmMessage string) *BtnElement {
	btn := NewBtn().Href(href).Title("Delete").Class("btn-outline-danger btn-sm px-1").
		Label(Icon("trash"))

	if confirmMessage == "" {
		confirmMessage = "Are you sure?"
	}

	btn.tag.AttributeUnsafe("onclick", "return confirm(\""+html.EscapeString(confirmMessage)+"\");")

	return btn
}

func NewIconBtn(href, icon string, label any) *BtnElement {
	return NewBtn().Href(href).Label(Icon(icon).Label(label))
}

// =================== Buttons panel ====================
type BtnPanelElement struct {
	buttons []*BtnElement
	classes dhtml.Classes
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*BtnPanelElement)(nil)

func NewBtnPanel() *BtnPanelElement {
	return &BtnPanelElement{}
}

func (e *BtnPanelElement) Class(v any) *BtnPanelElement {
	e.classes.Add(v)
	return e
}

func (e *BtnPanelElement) AddBtn(btn *BtnElement) *BtnPanelElement {
	e.buttons = append(e.buttons, btn)
	return e
}

func (e *BtnPanelElement) AddIconBtn(href, icon, label string) *BtnPanelElement {
	btn := NewBtn().Href(href).Label(Icon(icon).Label(label))

	e.buttons = append(e.buttons, btn)
	return e
}

func (e *BtnPanelElement) GetTags() dhtml.TagsList {
	body := dhtml.NewHtmlPiece()

	if len(e.buttons) == 0 {
		body.Append(dhtml.EmptyLabel("no buttons added"))
	} else {
		for _, btn := range e.buttons {
			body.Append(btn)
		}
	}

	rootTag := dhtml.Div().Class("border p-3").Class(e.classes).Append(body)

	return rootTag.GetTags()
}
