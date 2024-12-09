package mtweb

import (
	"github.com/mitoteam/dhtml"
)

func init() {
	dhtml.SetDefaultSubmitButtonClasses("btn btn-secondary")

	dhtml.FormManager.SetRenderErrorsF(func(fd *dhtml.FormData) (out dhtml.HtmlPiece) {
		container := dhtml.Div().Class("border p-3 border-danger mb-3")

		for name, itemErrors := range fd.GetErrors() {
			if container.HasChildren() {
				container.Append(dhtml.Div().Class("my-3 border"))
			}

			for _, itemError := range itemErrors {
				errorOut := dhtml.Div().Class("item-error")

				errorOut.Append(Icon("circle-xmark").Class("text-danger").ElementClass("me-1"))

				if name != "" {
					errorOut.Attribute("data-form-item-name", name).
						Append(dhtml.Span().Class("fw-bold").Text(name)).
						Append(":")
				}

				errorOut.Append(itemError)

				container.Append(errorOut)
			}
		}

		out.Append(container)
		return out
	})
}
