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

func BuildExperimentForm() *dhtml.FormElement {
	dhtml.FormManager.Register("test_form", &dhtml.FormHandler{
		RenderF: func() *dhtml.FormElement {
			input := dhtml.NewFormInput("text")
			control := dhtml.NewFormControl("test_name").Label("test label").Element(input)

			return dhtml.NewForm().
				Append(dhtml.Div().Text("test bdy").Class("border mb-3")).
				Append(control)
		},
	})

	return dhtml.FormManager.Render("test_form")
}
