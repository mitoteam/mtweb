package mtweb

import (
	"log"
	"slices"

	"github.com/mitoteam/dhtml"
)

// Font Awesome icons helper
// https://fontawesome.com

var (
	fa_icon_classes = []string{
		"fas", "far", "fal", "fad",
	}

	default_fa_class = "fas"
)

func SetDefaultFAClass(class string) {
	if slices.Contains(fa_icon_classes, class) {
		default_fa_class = class
	} else {
		log.Panicf("Unknown FontAwesome class: %s", class)
	}
}

type (
	IconElement struct {
		Name  string
		title string //always applied to icon itself
		label dhtml.HtmlPiece
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*IconElement)(nil)

func Icon(icon_name string) *IconElement {
	return &IconElement{
		Name: icon_name,
	}
}

func (i *IconElement) Title(s string) *IconElement {
	i.title = s
	return i
}

func (i *IconElement) Label(v any) *IconElement {
	i.label.Append(v)
	return i
}

func (i *IconElement) GetTags() dhtml.TagsList {
	icon_tag := dhtml.NewTag("i").Title(i.title)

	icon_tag.GetClasses().
		Add("fa-"+i.Name).
		AddFromSet(default_fa_class, fa_icon_classes)

	if i.label.IsEmpty() {
		//no label given, render just icon
		return icon_tag.GetTags()
	} else {
		//there is a label, so wrap everything in span
		icon_tag.Class("me-1")

		return dhtml.Span().
			Append(icon_tag).
			Append(i.label).
			GetTags()
	}
}
