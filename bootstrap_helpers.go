package mtweb

import "github.com/mitoteam/dhtml"

//Some Bootstrap Css Framework helpers
//https://getbootstrap.com/docs/5.3/getting-started/introduction/

func Card() dhtml.ElementI {
	return dhtml.NewTag("div").Class("card")
}
