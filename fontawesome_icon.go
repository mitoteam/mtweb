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
		"far", "fas", "fal", "fad",
	}

	default_fa_class = "far"
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
		Name           string
		title          string //always applied to icon itself
		label          dhtml.HtmlPiece
		iconClasses    dhtml.Classes
		elementClasses dhtml.Classes
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

// Classes applied to icon itself
func (i *IconElement) Class(v any) *IconElement {
	i.iconClasses.Add(v)
	return i
}

func (i *IconElement) ElementClass(v any) *IconElement {
	i.elementClasses.Add(v)
	return i
}

func (i *IconElement) GetTags() dhtml.TagsList {
	icon_tag := dhtml.NewTag("i").Title(i.title)

	icon_tag.GetClasses().
		Add("fa-"+i.Name).
		AddFromSet(default_fa_class, fa_icon_classes).
		Add(i.iconClasses)

	if i.label.IsEmpty() {
		//no label given, render just icon
		icon_tag.Class(i.elementClasses) // add container classes to icon itself
		return icon_tag.GetTags()
	} else {
		//there is a label, so wrap everything in span
		icon_tag.Class("me-1")

		return dhtml.Span().Class(i.elementClasses).
			Append(icon_tag).
			Append(i.label).
			GetTags()
	}
}
