package mtweb

import "github.com/mitoteam/dhtml"

//Some Font Awesome helpers
//https://fontawesome.com

type (
	Icon struct {
		Name  string
		Title string
		Label string
	}
)

// force interface implementation declaring fake variable
var _ dhtml.ElementI = (*Icon)(nil)

func (i *Icon) GetTags() dhtml.TagsList {
	r := make(dhtml.TagsList, 0, 2) //place for label without additional memory allocation

	t := dhtml.NewTag("i").
		Classes([]string{"fas", "fa-" + i.Name})

	if i.Title != "" {
		t.Attribute("title", i.Title)
	}

	r = append(r, t)

	if i.Label != "" {
		r = append(r, dhtml.Text(i.Label))
	}

	return r
}
