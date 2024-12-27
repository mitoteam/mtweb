package mtweb

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlbs"
	"github.com/mitoteam/dhtmlform"
)

// ======== Btn helpers ==========

func NewEditBtn(href string) *dhtmlbs.BtnElement {
	return dhtmlbs.NewBtn().Href(href).Title("Edit").Class("btn-sm px-1").
		Label(Icon("edit"))
}

func NewDeleteBtn(href, confirmMessage string) *dhtmlbs.BtnElement {
	btn := dhtmlbs.NewBtn().Href(href).Title("Delete").Class("btn-outline-danger btn-sm px-1").
		Label(Icon("trash"))

	if confirmMessage == "" {
		confirmMessage = "Are you sure?"
	}

	btn.Confirm(confirmMessage)

	return btn
}

func NewIconBtn(href, icon string, label any) *dhtmlbs.BtnElement {
	return dhtmlbs.NewBtn().Href(href).Label(Icon(icon).Label(label))
}

func NewIconSubmitBtn(icon, label string) *dhtmlform.FormControlElement {
	return dhtmlbs.NewSubmitBtn().Label(Icon(icon).Label(label))
}

func NewDefaultSubmitBtn() *dhtmlform.FormControlElement {
	return NewIconSubmitBtn("octagon-check", "Save")
	// return dhtmlbs.NewSubmitBtn().Label(mtweb.Icon("arrow-right-to-bracket").Label("Sign In"))
}

// =================== Buttons panel ====================
type BtnPanelElement struct {
	buttons []*dhtmlbs.BtnElement
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

func (e *BtnPanelElement) AddBtn(btn *dhtmlbs.BtnElement) *BtnPanelElement {
	e.buttons = append(e.buttons, btn)
	return e
}

func (e *BtnPanelElement) AddIconBtn(href, icon, label string) *BtnPanelElement {
	btn := dhtmlbs.NewBtn().Href(href).Label(Icon(icon).Label(label))

	e.buttons = append(e.buttons, btn)
	return e
}

func (e *BtnPanelElement) GetTags() dhtml.TagList {
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
