package gtml

import (
	"context"
	"github.com/a-h/templ"
	"io"
)

type (
	// Attrs is a map of attributes for the element, currently uses templ.Attributes.
	Attrs templ.Attributes
	// ContentType represents the type of content in an element.
	ContentType int
	// Element represents an HTML element.
	Element struct {
		Attrs       Attrs             `json:"attrs"`
		Children    []templ.Component `json:"children"`
		Tag         string            `json:"tag"`
		Content     string            `json:"content"`
		ContentType ContentType       `json:"content_type"`
	}
)

const (
	// TextContent represents text content to be escaped.
	TextContent ContentType = iota
	// RawContent represents raw HTML content with no escaping.
	RawContent
)

var (
	// VoidElements is a list of HTML void elements, these don't have children or content.
	VoidElements = map[string]string{
		"area":   ">",
		"base":   ">",
		"br":     ">",
		"col":    ">",
		"embed":  ">",
		"hr":     ">",
		"img":    ">",
		"input":  ">",
		"link":   ">",
		"meta":   ">",
		"source": ">",
		"track":  ">",
		"wbr":    ">",
	}
	// NA is an empty set of attributes.
	NA = Attrs{}
)

// A creates an anchor element <a></a>.
func A(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "a", Attrs: attrs, Children: children}
}

// Address creates an address element <address></address>.
func Address(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "address", Attrs: attrs, Children: children}
}

// Article creates an article element <article></article>.
func Article(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "article", Attrs: attrs, Children: children}
}

// Aside creates an aside element <aside></aside>.
func Aside(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "aside", Attrs: attrs, Children: children}
}

// B creates a bold element <b></b>.
func B(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "b", Attrs: attrs, Children: children}
}

// Blockquote creates a blockquote element <blockquote></blockquote>.
func Blockquote(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "blockquote", Attrs: attrs, Children: children}
}

// Body creates a body element <body></body>.
func Body(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "body", Attrs: attrs, Children: children}
}

// Button creates a button element <button></button>.
func Button(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "button", Attrs: attrs, Children: children}
}

// BR creates a line break element <br />.
func BR(attrs Attrs) *Element {
	return &Element{Tag: "br", Attrs: attrs, Children: []templ.Component{}}
}

// Caption creates a caption element <caption></caption>.
func Caption(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "caption", Attrs: attrs, Children: children}
}

// DD creates a definition description element <dd></dd>.
func DD(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "dd", Attrs: attrs, Children: children}
}

// Details creates a details element <details></details>.
func Details(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "details", Attrs: attrs, Children: children}
}

// Dialog creates a dialog element <dialog></dialog>.
func Dialog(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "dialog", Attrs: attrs, Children: children}
}

// Div creates a div element <div></div>.
func Div(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "div", Attrs: attrs, Children: children}
}

// DL creates a definition list element <dl></dl>.
func DL(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "dl", Attrs: attrs, Children: children}
}

// DT creates a definition term element <dt></dt>.
func DT(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "dt", Attrs: attrs, Children: children}
}

// Empty creates an empty element.
func Empty() *Element {
	return &Element{}
}

// Fieldset creates a fieldset element <fieldset></fieldset>.
func Fieldset(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "fieldset", Attrs: attrs, Children: children}
}

// Form creates a form element <form></form>.
func Form(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "form", Attrs: attrs, Children: children}
}

// Fragment allows you to create a fragment of elements without a parent element.
func Fragment(children ...templ.Component) *Element {
	return &Element{Children: children}
}

// H1 creates a heading level 1 element <h1></h1>.
func H1(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "h1", Attrs: attrs, Children: children}
}

// H2 creates a heading level 2 element <h2></h2>.
func H2(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "h2", Attrs: attrs, Children: children}
}

// H3 creates a heading level 3 element <h3></h3>.
func H3(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "h3", Attrs: attrs, Children: children}
}

// H4 creates a heading level 4 element <h4></h4>.
func H4(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "h4", Attrs: attrs, Children: children}
}

// H5 creates a heading level 5 element <h5></h5>.
func H5(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "h5", Attrs: attrs, Children: children}
}

// Head creates a head element <head></head>.
func Head(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "head", Attrs: attrs, Children: children}
}

// Header creates a header element <header></header>.
func Header(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "header", Attrs: attrs, Children: children}
}

// HR creates a horizontal rule element <hr />.
func HR(attrs Attrs) *Element {
	return &Element{Tag: "hr", Attrs: attrs, Children: []templ.Component{}}
}

// Html creates an html element <html></html>.
func Html(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "html", Attrs: attrs, Children: children}
}

// I creates an italic element <i></i>.
func I(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "i", Attrs: attrs, Children: children}
}

// Img creates an image element <img />.
func Img(attrs Attrs) *Element {
	return &Element{Tag: "img", Attrs: attrs, Children: []templ.Component{}}
}

// Input creates an input element <input />.
func Input(attrs Attrs) *Element {
	return &Element{Tag: "input", Attrs: attrs, Children: []templ.Component{}}
}

// Label creates a label element <label></label>.
func Label(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "label", Attrs: attrs, Children: children}
}

// Legend creates a legend element <legend></legend>.
func Legend(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "legend", Attrs: attrs, Children: children}
}

// LI creates a list item element <li></li>.
func LI(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "li", Attrs: attrs, Children: children}
}

// Link creates a link element <link />.
func Link(attrs Attrs) *Element {
	return &Element{Tag: "link", Attrs: attrs, Children: []templ.Component{}}
}

// Main creates a main element <main></main>.
func Main(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "main", Attrs: attrs, Children: children}
}

// Meta creates a meta element <meta />.
func Meta(attrs Attrs) *Element {
	return &Element{Tag: "meta", Attrs: attrs, Children: []templ.Component{}}
}

// Nav creates a nav element <nav></nav>.
func Nav(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "nav", Attrs: attrs, Children: children}
}

// OL creates an ordered list element <ol></ol>.
func OL(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "ol", Attrs: attrs, Children: children}
}

// Optgroup creates an option group element <optgroup></optgroup>.
func Optgroup(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "optgroup", Attrs: attrs, Children: children}
}

// Option creates an option element <option></option>.
func Option(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "option", Attrs: attrs, Children: children}
}

// P creates a paragraph element <p></p>.
func P(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "p", Attrs: attrs, Children: children}
}

// Raw creates an element with raw HTML content.
// Warning: this elements cannot have children and content is not escaped.
func Raw(html string) *Element {
	return &Element{Content: html, ContentType: RawContent, Children: []templ.Component{}}
}

// Section creates a section element <section></section>.
func Section(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "section", Attrs: attrs, Children: children}
}

// Select creates a select element <select></select>.
func Select(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "select", Attrs: attrs, Children: children}
}

// Small creates a small element <small></small>.
func Small(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "small", Attrs: attrs, Children: children}
}

// Span creates a span element <span></span>.
func Span(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "span", Attrs: attrs, Children: children}
}

// Strong creates a strong element <strong></strong>.
func Strong(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "strong", Attrs: attrs, Children: children}
}

// Summary creates a summary element <summary></summary>.
func Summary(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "summary", Attrs: attrs, Children: children}
}

// Table creates a table element <table></table>.
func Table(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "table", Attrs: attrs, Children: children}
}

// TBody creates a table body element <tbody></tbody>.
func TBody(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "tbody", Attrs: attrs, Children: children}
}

// TD creates a table data element <td></td>.
func TD(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "td", Attrs: attrs, Children: children}
}

// Text creates an element with text content.
// Warning: this elements cannot have children and content will be escaped.
func Text(text string) *Element {
	return &Element{Content: text, ContentType: TextContent, Children: []templ.Component{}}
}

// Textarea creates a textarea element <textarea></textarea>.
func Textarea(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "textarea", Attrs: attrs, Children: children}
}

// TFoot creates a table footer element <tfoot></tfoot>.
func TFoot(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "tfoot", Attrs: attrs, Children: children}
}

// TH creates a table header element <th></th>.
func TH(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "th", Attrs: attrs, Children: children}
}

// THead creates a table header element <thead></thead>.
func THead(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "thead", Attrs: attrs, Children: children}
}

// Title creates a title element <title></title>.
func Title(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "title", Attrs: attrs, Children: children}
}

// TR creates a table row element <tr></tr>.
func TR(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "tr", Attrs: attrs, Children: children}
}

// UL creates an unordered list element <ul></ul>.
func UL(attrs Attrs, children ...templ.Component) *Element {
	return &Element{Tag: "ul", Attrs: attrs, Children: children}
}

// AddChildren adds children *Element's to this element.
func (e *Element) AddChildren(children ...templ.Component) {
	e.Children = append(e.Children, children...)
}

// Render renders the Element to the writer.
func (e *Element) Render(ctx context.Context, w io.Writer) error {
	closer := ">"
	voidCloser, isVoid := VoidElements[e.Tag]
	if isVoid {
		closer = voidCloser
	}

	if e.Tag != "" {
		if _, err := io.WriteString(w, "<"+e.Tag); err != nil {
			return err
		}
		if len(e.Attrs) > 0 {
			if err := templ.RenderAttributes(ctx, w, templ.Attributes(e.Attrs)); err != nil {
				return err
			}
		}
		if _, err := io.WriteString(w, closer); err != nil {
			return err
		}
	}

	if !isVoid {
		if e.Content != "" {
			switch e.ContentType {
			case TextContent:
				if _, err := io.WriteString(w, templ.EscapeString(e.Content)); err != nil {
					return err
				}
			case RawContent:
				if _, err := io.WriteString(w, e.Content); err != nil {
					return err
				}
			}
		} else {
			for _, child := range e.Children {
				if err := child.Render(ctx, w); err != nil {
					return err
				}
			}
		}

		if e.Tag != "" {
			if _, err := io.WriteString(w, "</"+e.Tag+">"); err != nil {
				return err
			}
		}
	}

	return nil
}
