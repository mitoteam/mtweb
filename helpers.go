package mtweb

import "github.com/mitoteam/dhtml"

func NewTable() *dhtml.TableElement {
	return dhtml.NewTable().Class("table table-hover table-sm").EmptyLabel("nothing here yet")
}
