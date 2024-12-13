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
	DefaultBtnSizeClass = BtnSizeClasses[0]

	BtnVariantClasses = []string{
		"btn-outline-primary", "btn-outline-secondary", "btn-outline-success", "btn-outline-danger", "btn-outline-warning",
		"btn-outline-info", "btn-outline-light", "btn-outline-dark",
		"btn-primary", "btn-secondary", "btn-success", "btn-danger", "btn-warning", "btn-info", "btn-light", "btn-dark",
	}
	DefaultBtnVariantClass = BtnVariantClasses[0]
)

type Btn struct {
	tag *dhtml.Tag
}

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*Btn)(nil)

func NewBtn() *Btn {
	return &Btn{
		tag: dhtml.NewTag("a"),
	}
}

func (e *Btn) Href(url string) *Btn {
	e.tag.Attribute("href", url)
	return e
}

func (e *Btn) Label(v any) *Btn {
	e.tag.Append(v)
	return e
}

func (e *Btn) Title(title string) *Btn {
	e.tag.Title(title)
	return e
}

func (e *Btn) Class(v any) *Btn {
	e.tag.Class(v)
	return e
}

func (e *Btn) GetTags() dhtml.TagsList {
	e.tag.GetClasses().
		Add("btn text-nowrap").
		AddFromSet(DefaultBtnVariantClass, BtnVariantClasses).
		AddFromSet(DefaultBtnSizeClass, BtnSizeClasses)

	return e.tag.GetTags()
}

// ==================
func NewEditBtn(href string) *Btn {
	return NewBtn().Href(href).Title("Edit").Class("btn-sm px-1").
		Label(Icon("edit"))
}

func NewDeleteBtn(href, confirmMessage string) *Btn {
	btn := NewBtn().Href(href).Title("Delete").Class("btn-outline-danger btn-sm px-1").
		Label(Icon("trash"))

	if confirmMessage == "" {
		confirmMessage = "Are you sure?"
	}

	btn.tag.AttributeUnsafe("onclick", "return confirm(\""+html.EscapeString(confirmMessage)+"\");")

	return btn
}
