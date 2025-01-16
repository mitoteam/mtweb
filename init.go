package mtweb

import (
	"github.com/mitoteam/dhtml"
)

func init() {
	dhtml.Settings().EmptyLabelRendererF = func(label string, span *dhtml.Tag) {
		if label == "" {
			label = "empty"
		}

		span.Append("[" + label + "]").Class("text-muted")
	}
}
