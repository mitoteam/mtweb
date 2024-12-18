package mtweb

import (
	"time"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
)

func BuildExperimentHtml() string {
	document := dhtml.NewHtmlDocument()

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

var ExperimentFormHandler = dhtmlform.FormHandler{
	RenderF: func(formBody *dhtml.HtmlPiece, fd *dhtmlform.FormData) {
		formBody.Append("ExperimentFormHandler")
		formBody.Append(dhtmlform.NewTextarea("area").Default("def value\nmulti").Label("Label").Note("notes for <textarea>"))

		formBody.Append(
			dhtml.Div().Append("Deeper").Class("mt-3 p-3 border").Append(
				dhtmlform.NewTextarea("area2").Default("def2").Label("Label2").Note("note2"),
			).Append(
				dhtml.Div().Append("And Deeper").Class("mt-3 p-3 border").Append(
					dhtmlform.NewTextarea("area3").Default("def3").Label("Label3").Note("note3"),
				),
			),
		)

		formBody.Append(dhtmlform.NewSubmitBtn())
	},
	ValidateF: func(fd *dhtmlform.FormData) {
		fd.SetRebuild(true)
	},
}

func init() {
	dhtml.FormManager.Register(&dhtml.FormHandler{
		Id: "test_form",
		RenderF: func(form *dhtml.FormElement, fd *dhtml.FormData) {
			//form.Append(dhtml.Dbg("Args: %v", fd.GetAllArgs()))

			form.
				Append(
					dhtml.NewFormInput("weha", "text").
						Label("test label").Note("test note").DefaultValue("default v"),
				).
				Append(
					dhtml.Div().Class("p-3 border").Append(
						dhtml.NewFormInput("another", "date").
							Label("Date").Note("Another input").DefaultValue(time.Now().Format(time.DateOnly)),
					),
				).
				Append(dhtml.NewFormSubmit())
		},
		ValidateF: func(fd *dhtml.FormData) {
			if v, ok := fd.GetValue("weha").(string); ok {
				if len(v) < 3 {
					fd.SetItemError("weha", "At least three characters expected")
				}
			}
		},
		SubmitF: func(fd *dhtml.FormData) {
			fd.SetRedirect("/")
		},
	})
}
