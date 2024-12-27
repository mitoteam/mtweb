package mtweb

import (
	"log"
	"slices"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/mttools"
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

func (i *IconElement) GetTags() dhtml.TagList {
	icon_tag := dhtml.NewTag("i").Title(i.title)

	icon_tag.GetClasses().
		Add("fa-"+i.Name).
		AddFromSet(fa_icon_classes, default_fa_class).
		Add(i.iconClasses)

	if i.label.IsEmpty() {
		//no label given, render just icon
		icon_tag.Class(i.elementClasses) // add container classes to icon itself
		return icon_tag.GetTags()
	} else {
		//there is a label, so wrap everything in span
		icon_tag.Class("me-2")

		return dhtml.Span().Class(i.elementClasses).
			Append(icon_tag).
			Append(i.label).
			GetTags()
	}
}

// ================== Useful helpers ==========================

func IconYes() *IconElement {
	return Icon("check")
}

func IconNo() *IconElement {
	return Icon("ban")
}

func IconYesNo(v any) *IconElement {
	if mttools.IsEmpty(v) {
		return IconNo()
	} else {
		return IconYes()
	}
}

func IconYesNoTitle(v any, yesTitle, noTitle string) *IconElement {
	if mttools.IsEmpty(v) {
		return IconNo().Title(yesTitle)
	} else {
		return IconYes().Title(noTitle)
	}
}

// Yes/No icon with label.
// First arg is value itself.
// Second arg - label is for YES value
// Third arg - label is for NO value
// Fourth arg - title is for YES icon
// Fifth arg - title is for NO icon
func IconYesNoLabel(v any, labels ...any) *IconElement {
	if mttools.IsEmpty(v) {
		icon := IconNo()

		if len(labels) > 0 {
			icon.Label(labels[0])
		}

		if len(labels) > 2 {
			icon.Label(labels[2])
		}

		return icon
	} else {
		icon := IconYes()

		if len(labels) > 1 {
			icon.Label(labels[1])
		}

		if len(labels) > 3 {
			icon.Label(labels[3])
		}

		return icon
	}
}
