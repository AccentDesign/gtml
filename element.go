package gtml

import (
	"context"
	"github.com/a-h/templ"
	"io"
	"slices"
	"strings"
)

type (
	Attrs   templ.Attributes
	Element struct {
		Attrs    Attrs      `json:"attrs"`
		Children []*Element `json:"children"`
		Tag      string     `json:"tag"`
		Text     string     `json:"text"`
	}
)

var (
	VoidElements = []string{
		"area",
		"base",
		"br",
		"col",
		"embed",
		"hr",
		"img",
		"input",
		"link",
		"meta",
		"source",
		"track",
		"wbr",
	}
	NA = Attrs{}
)

func A(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "a", Attrs: attrs, Children: children}
}

func Address(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "address", Attrs: attrs, Children: children}
}

func Article(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "article", Attrs: attrs, Children: children}
}

func Aside(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "aside", Attrs: attrs, Children: children}
}

func B(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "b", Attrs: attrs, Children: children}
}

func Blockquote(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "blockquote", Attrs: attrs, Children: children}
}

func Body(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "body", Attrs: attrs, Children: children}
}

func Button(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "button", Attrs: attrs, Children: children}
}

func BR(attrs Attrs) *Element {
	return &Element{Tag: "br", Attrs: attrs, Children: []*Element{}}
}

func Caption(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "caption", Attrs: attrs, Children: children}
}

func DD(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "dd", Attrs: attrs, Children: children}
}

func Details(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "details", Attrs: attrs, Children: children}
}

func Dialog(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "dialog", Attrs: attrs, Children: children}
}

func Div(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "div", Attrs: attrs, Children: children}
}

func DL(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "dl", Attrs: attrs, Children: children}
}

func DT(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "dt", Attrs: attrs, Children: children}
}

func Fieldset(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "fieldset", Attrs: attrs, Children: children}
}

func Form(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "form", Attrs: attrs, Children: children}
}

func H1(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "h1", Attrs: attrs, Children: children}
}

func H2(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "h2", Attrs: attrs, Children: children}
}

func H3(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "h3", Attrs: attrs, Children: children}
}

func H4(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "h4", Attrs: attrs, Children: children}
}

func H5(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "h5", Attrs: attrs, Children: children}
}

func Header(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "header", Attrs: attrs, Children: children}
}

func HR(attrs Attrs) *Element {
	return &Element{Tag: "hr", Attrs: attrs, Children: []*Element{}}
}

func I(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "i", Attrs: attrs, Children: children}
}

func Img(attrs Attrs) *Element {
	return &Element{Tag: "img", Attrs: attrs, Children: []*Element{}}
}

func Input(attrs Attrs) *Element {
	return &Element{Tag: "input", Attrs: attrs, Children: []*Element{}}
}

func Label(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "label", Attrs: attrs, Children: children}
}

func Legend(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "legend", Attrs: attrs, Children: children}
}

func LI(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "li", Attrs: attrs, Children: children}
}

func Main(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "main", Attrs: attrs, Children: children}
}

func Nav(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "nav", Attrs: attrs, Children: children}
}

func OL(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "ol", Attrs: attrs, Children: children}
}

func Optgroup(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "optgroup", Attrs: attrs, Children: children}
}

func Option(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "option", Attrs: attrs, Children: children}
}

func P(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "p", Attrs: attrs, Children: children}
}

func Section(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "section", Attrs: attrs, Children: children}
}

func Select(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "select", Attrs: attrs, Children: children}
}

func Small(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "small", Attrs: attrs, Children: children}
}

func Span(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "span", Attrs: attrs, Children: children}
}

func Strong(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "strong", Attrs: attrs, Children: children}
}

func Summary(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "summary", Attrs: attrs, Children: children}
}

func Table(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "table", Attrs: attrs, Children: children}
}

func TBody(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "tbody", Attrs: attrs, Children: children}
}

func TD(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "td", Attrs: attrs, Children: children}
}

func Text(text string) *Element {
	return &Element{Text: text, Children: []*Element{}}
}

func Textarea(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "textarea", Attrs: attrs, Children: children}
}

func TFoot(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "tfoot", Attrs: attrs, Children: children}
}

func TH(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "th", Attrs: attrs, Children: children}
}

func THead(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "thead", Attrs: attrs, Children: children}
}

func TR(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "tr", Attrs: attrs, Children: children}
}

func UL(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "ul", Attrs: attrs, Children: children}
}

func (e *Element) AddChildren(children ...*Element) {
	e.Children = append(e.Children, children...)
}

func (e *Element) Render(ctx context.Context, w io.Writer) error {
	var b strings.Builder

	isVoid := slices.Contains(VoidElements, e.Tag)

	if e.Tag != "" {
		b.WriteString("<" + e.Tag)

		if len(e.Attrs) > 0 {
			if err := templ.RenderAttributes(ctx, &b, templ.Attributes(e.Attrs)); err != nil {
				return err
			}
		}

		if isVoid {
			b.WriteString(" />")
		} else {
			b.WriteString(">")
		}
	}

	if !isVoid {
		if e.Text != "" {
			b.WriteString(templ.EscapeString(e.Text))
		} else {
			for _, child := range e.Children {
				if err := child.Render(ctx, &b); err != nil {
					return err
				}
			}
		}

		if e.Tag != "" {
			b.WriteString("</" + e.Tag + ">")
		}
	}

	_, err := io.WriteString(w, b.String())
	return err
}
