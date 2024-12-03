package mtweb

import (
	"github.com/mitoteam/dhtml"
)

func BuildExperimentHtml() *dhtml.Tag {
	head := dhtml.NewTag("head").
		Append(
			dhtml.NewTag("link").
				Attribute("href", "/assets/vendor/bootstrap.min.css").Attribute("rel", "stylesheet"),
		).
		Append(
			dhtml.NewTag("link").
				Attribute("href", "/assets/vendor/fontawesome.min.css").Attribute("rel", "stylesheet"),
		).
		Append(
			dhtml.NewTag("link").
				Attribute("href", "/assets/vendor/regular.min.css").Attribute("rel", "stylesheet"),
		)

	body := dhtml.NewTag("body")

	div := dhtml.Div().
		Id("test").
		Attribute("data-mt-test", "some attribute").
		//classes deduplication
		Class("border").Class("m-3").Class("p-3").Class("border").
		Content("some <escaped> text")

	body.Append(div)

	body.
		Append(
			dhtml.Div().Class("border").Class("t-3").Class("p-3").
				Content("multi").
				Append(
					dhtml.Span().Content("red").Classes([]string{"border", "border-danger", "border-5"}),
				).
				Content("content"),
		).
		Append(
			dhtml.Div().Classes([]string{"border", "p-3", "m-3"}).
				Content("content").
				Content("only"),
		).
		Append(
			dhtml.Div().Classes([]string{"border", "p-3", "m-3"}).
				Content("Icon test: ").
				Append(&Icon{Name: "face-awesome", Label: "test"}),
		)

	html := dhtml.NewTag("html").
		Comment("some <escaped> comment").
		Append(head).
		Append(body)

	return html
}
