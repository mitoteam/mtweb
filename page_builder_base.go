package mtweb

import (
	"fmt"

	"github.com/mitoteam/dhtml"
	"github.com/mitoteam/dhtmlform"
	"github.com/mitoteam/mbr"
)

type PageBuilderBase struct {
	Ctx     *mbr.MbrContext
	Regions dhtml.NamedHtmlPieces

	RenderF func() error

	title           string // page title (usually rendered as <H1> tag)
	BuildHeadTitleF func() string

	document *dhtml.HtmlDocument
}

func NewPageBuilderBase(ctx *mbr.MbrContext) PageBuilderBase {
	p := PageBuilderBase{
		Ctx:     ctx,
		Regions: dhtml.NewNamedHtmlPieces(),
	}

	//default title builder
	p.BuildHeadTitleF = func() string {
		return p.title // just page title by default
	}

	//create document
	p.document = dhtml.NewHtmlDocument()

	//default page settings and assets
	p.document.
		Charset("utf-8").
		Stylesheet("/assets/mtweb/vendor/bootstrap.min.css").
		Stylesheet("/assets/mtweb/vendor/fontawesome.min.css").
		Stylesheet("/assets/mtweb/vendor/regular.min.css")

	return p
}

// Renders the page to HTML string
func (p *PageBuilderBase) Render() (string, error) {
	p.Ctx.Writer().Header().Add("Content-Type", "text/html;charset=utf-8")

	p.document.Title(p.BuildHeadTitleF())

	if p.RenderF == nil {
		return "", fmt.Errorf("PageBuilderBase.RenderF is not set")
	} else {
		//JS scripts from mtweb assets
		p.document.Body().
			Append(dhtml.NewTag("script").Attribute("src", "/assets/mtweb/vendor/bootstrap.bundle.min.js"))

		// render page content
		if err := p.RenderF(); err != nil {
			return "", err
		}

		return dhtml.Piece(p.document).String(), nil
	}
}

// very root html document element, to be used for rendering page content
func (p *PageBuilderBase) GetDocument() *dhtml.HtmlDocument {
	return p.document
}

// Sets page's title
func (p *PageBuilderBase) Title(title string) *PageBuilderBase {
	p.title = title
	return p
}

func (p *PageBuilderBase) GetTitle() string {
	return p.title
}

// by Default every page builder has at least one region called "main"
func (p *PageBuilderBase) Main(v any) *PageBuilderBase {
	p.Regions.Add("main", v)
	return p
}

// contents of the "main" region
func (p *PageBuilderBase) GetMain() *dhtml.HtmlPiece {
	return p.Regions.Get("main")
}

// Builds new dhtml.FormContext to be used with form builder
func (p *PageBuilderBase) NewFormContext() *dhtmlform.FormContext {
	fc := dhtmlform.NewFormContext(p.Ctx.Writer(), p.Ctx.Request())

	// some useful for every form things
	fc.SetParam("MbrContext", p.Ctx)

	//default redirect from "destination" query parameter
	if destination := p.Ctx.Request().URL.Query().Get("destination"); destination != "" {
		fc.SetRedirect(destination)
	}

	return fc
}
