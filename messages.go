package mtweb

import (
	"fmt"

	"github.com/mitoteam/dhtml"
)

func RenderInfo(message any) *dhtml.HtmlPiece {
	out := dhtml.Div().Class("text-secondary").
		Append(Icon(FaIconInfo).Label(message))

	return dhtml.Piece(out)
}

func RenderInfof(message string, args ...any) *dhtml.HtmlPiece {
	return RenderInfo(fmt.Sprintf(message, args...))
}

func RenderWarning(message any) *dhtml.HtmlPiece {
	out := dhtml.Div().Class("text-warning").
		Append(Icon(FaIconWarning).Label(message))

	return dhtml.Piece(out)
}

func RenderWarningf(message string, args ...any) *dhtml.HtmlPiece {
	return RenderWarning(fmt.Sprintf(message, args...))
}

func RenderError(message any) *dhtml.HtmlPiece {
	out := dhtml.Div().Class("text-danger").
		Append(Icon(FaIconError).Label(message))

	return dhtml.Piece(out)
}

func RenderErrorf(message string, args ...any) *dhtml.HtmlPiece {
	return RenderError(fmt.Sprintf(message, args...))
}
