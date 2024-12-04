package mtweb

import (
	"github.com/mitoteam/dhtml"
)

func BuildExperimentHtml() string {
	document := dhtml.NewDocument()

	document.
		Title("The Experiment!").
		Stylesheet("/assets/vendor/bootstrap.min.css").
		Stylesheet("/assets/vendor/fontawesome.min.css").
		Stylesheet("/assets/vendor/regular.min.css")

	div := dhtml.Div().
		Id("test").
		Attribute("data-mt-test", "some attribute").
		//classes deduplication
		Class("border").Class("m-3").Class("p-3").Class("border").
		Text("some <escaped> text")

	document.Body().Append(div)

	document.Body().
		Append(
			dhtml.Div().Class("border").Class("t-3").Class("p-3").
				Text("multi").
				Append(
					dhtml.Span().Text("red").Class([]string{"border", "border-danger", "border-5"}),
				).
				Text("content"),
		).
		Append(
			dhtml.Div().Class("border p-3   m-3").
				Text("content").
				Text("only"),
		).
		Append(
			dhtml.Div().Class([]string{"border", "p-3", "m-3"}).
				Text("Icon test: ").
				Append(Icon("face-awesome").Label("label test")),
		)

	document.Body().Append(
		NewCard().
			Header(
				NewJustifiedLR().
					L("Card title text").
					R(Icon("car").Title("car icon title")),
			).
			Body("card body"),
	)

	return document.String()
}

func BuildExperimentForm() *dhtml.HtmlPiece {
	return dhtml.Piece("Form body")
}