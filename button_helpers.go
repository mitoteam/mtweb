package mtweb

import (
	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlbs"
	"github.com/mitoteam/dhtmlform"
	"github.com/mitoteam/mbr"
)

// ======== Btn helpers ==========

func NewEditBtn(href string) *dhtmlbs.BtnElement {
	return dhtmlbs.NewBtn().Href(href).Title("Edit").Class("btn-sm px-1").
		Label(Icon(FaIconEdit))
}

func NewDeleteBtn(href, confirmMessage string) *dhtmlbs.BtnElement {
	btn := dhtmlbs.NewBtn().Href(href).Title("Delete").Class("btn-outline-danger btn-sm px-1").
		Label(Icon(FaIconDelete))

	if confirmMessage == "" {
		confirmMessage = "Are you sure?"
	}

	btn.Confirm(confirmMessage)

	return btn
}

func NewIconBtn(href, icon string, label any) *dhtmlbs.BtnElement {
	return dhtmlbs.NewBtn().Href(href).Label(Icon(icon).Label(label))
}

// Same as NewIconBtn but with route args instead of raw href
func NewIconBtnR(icon string, label any, routeRef any, routeArgs ...any) *dhtmlbs.BtnElement {
	return NewIconBtn(mbr.Url(routeRef, routeArgs...), icon, label)
}

// small btm a-la "edit" but with custom icon
func NewSmBtn(href, icon string) *dhtmlbs.BtnElement {
	return dhtmlbs.NewBtn().Href(href).Class("btn-sm px-1").Label(Icon(icon))
}

// Same as NewSmBtn but with route args instead of raw href
func NewSmBtnR(icon string, routeRef any, routeArgs ...any) *dhtmlbs.BtnElement {
	return NewSmBtn(mbr.Url(routeRef, routeArgs...), icon)
}

func NewIconSubmitBtn(icon, label string) *dhtmlform.FormControlElement {
	return dhtmlbs.NewSubmitBtn().Label(Icon(icon).Label(label))
}

func NewDefaultSubmitBtn() *dhtmlform.FormControlElement {
	return NewIconSubmitBtn("octagon-check", "Save")
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

	rootTag := dhtml.Div().Class("border p-3 bg-light-subtle").Class(e.classes).Append(body)

	return rootTag.GetTags()
}
