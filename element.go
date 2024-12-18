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
		"br",
		"hr",
		"img",
		"input",
	}
	NA = Attrs{}
)

func A(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "a", Attrs: attrs, Children: children}
}

func Aside(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "aside", Attrs: attrs, Children: children}
}

func Button(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "button", Attrs: attrs, Children: children}
}

func BR(attrs Attrs) *Element {
	return &Element{Tag: "br", Attrs: attrs}
}

func Caption(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "caption", Attrs: attrs, Children: children}
}

func Details(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "details", Attrs: attrs, Children: children}
}

func Div(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "div", Attrs: attrs, Children: children}
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
	return &Element{Tag: "hr", Attrs: attrs}
}

func Img(attrs Attrs) *Element {
	return &Element{Tag: "img", Attrs: attrs}
}

func Input(attrs Attrs) *Element {
	return &Element{Tag: "input", Attrs: attrs}
}

func Label(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "label", Attrs: attrs, Children: children}
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

func Span(attrs Attrs, children ...*Element) *Element {
	return &Element{Tag: "span", Attrs: attrs, Children: children}
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
	return &Element{Text: text}
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

	if e.Tag != "" {
		b.WriteString("<" + e.Tag)

		if _, err := io.WriteString(w, b.String()); err != nil {
			return err
		}
		b.Reset()

		if len(e.Attrs) > 0 {
			if err := templ.RenderAttributes(ctx, w, templ.Attributes(e.Attrs)); err != nil {
				return err
			}
		}
		if slices.Contains(VoidElements, e.Tag) {
			b.WriteString(" />")
		} else {
			b.WriteString(">")
		}
	}

	if e.Text != "" {
		b.WriteString(e.Text)
	} else {
		for _, child := range e.Children {
			if err := child.Render(ctx, &b); err != nil {
				return err
			}
		}
	}

	if e.Tag != "" && !slices.Contains(VoidElements, e.Tag) {
		b.WriteString("</" + e.Tag + ">")
	}

	_, err := io.WriteString(w, b.String())
	return err
}
