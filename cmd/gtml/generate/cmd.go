package generate

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/accentdesign/gtml"
	"go/format"
	"golang.org/x/net/html"
	"io"
	"log/slog"
	"strings"
)

func Run(log *slog.Logger, r io.Reader, w2 io.Writer) error {
	var b bytes.Buffer
	w := &statefulWriter{w: &b}
	w.Write("package html\n")
	w.Write("\n")
	w.Write("import (\n")
	w.Write("\t. \"github.com/accentdesign/gtml\"\n")
	w.Write(")\n")
	w.Write("\n")
	w.Write("func Component() *Element {\n")
	w.Write("\treturn ")
	z := html.NewTokenizer(r)
	var hasContent bool
	var depth int

loop:
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if err := z.Err(); err != nil {
				if errors.Is(err, io.EOF) {
					if !hasContent {
						w.Write("Empty()")
					}
					break loop
				}
				return err
			}

		case html.TextToken:
			text := string(z.Text())
			trimmed := strings.TrimSpace(text)
			if trimmed == "" {
				continue
			}
			hasContent = true
			if depth > 0 {
				w.Write(",")
			}
			w.Write(fmt.Sprintf("Text(%q)", trimmed))

		case html.StartTagToken, html.SelfClosingTagToken:
			hasContent = true
			name, hasAttr := z.TagName()

			fn, isFn := funcMap[string(name)]

			if depth > 0 {
				w.Write(",")
			}

			if isFn {
				w.Write(fn)
				w.Write("(")
			} else {
				w.Write("New(")
				w.Write(fmt.Sprintf("%q", string(name)))
				w.Write(",")
			}

			if hasAttr {
				w.Write("Attrs{")
				for {
					key, val, moreAttr := z.TagAttr()
					w.Write("\"" + string(key) + "\":")
					if _, ok := boolAttrs[string(key)]; ok {
						w.Write("true")
					} else {
						w.Write("\"" + string(val) + "\"")
					}
					if moreAttr {
						w.Write(",")
					} else {
						w.Write("}")
						break
					}
				}
			} else {
				w.Write("NA")
			}

			depth++

			_, isVoid := gtml.VoidElements[string(name)]

			if tt == html.SelfClosingTagToken || isVoid {
				depth--
				w.Write(")")
			}

		case html.EndTagToken:
			depth--
			w.Write(")")
		}
	}

	w.Write("\n}\n")
	if w.err != nil {
		return w.err
	}

	formatted, err := format.Source(b.Bytes())
	if err != nil {
		return fmt.Errorf("error formatting output: %w", err)
	}
	if _, err = w2.Write(formatted); err != nil {
		return fmt.Errorf("error writing output: %w", err)
	}

	return nil
}

// statefulWriter only writes if no errors have occurred earlier in its lifetime.
type statefulWriter struct {
	w   io.Writer
	err error
}

func (w *statefulWriter) Write(s string) {
	if w.err != nil {
		return
	}
	_, w.err = w.w.Write([]byte(s))
}

var funcMap = map[string]string{
	"a":          "A",
	"address":    "Address",
	"article":    "Article",
	"aside":      "Aside",
	"b":          "B",
	"blockquote": "Blockquote",
	"body":       "Body",
	"button":     "Button",
	"br":         "BR",
	"caption":    "Caption",
	"dd":         "DD",
	"details":    "Details",
	"dialog":     "Dialog",
	"div":        "Div",
	"dl":         "DL",
	"dt":         "DT",
	"fieldset":   "Fieldset",
	"form":       "Form",
	"h1":         "H1",
	"h2":         "H2",
	"h3":         "H3",
	"h4":         "H4",
	"h5":         "H5",
	"head":       "Head",
	"header":     "Header",
	"hr":         "HR",
	"html":       "Html",
	"i":          "I",
	"img":        "Img",
	"input":      "Input",
	"label":      "Label",
	"legend":     "Legend",
	"li":         "LI",
	"link":       "Link",
	"main":       "Main",
	"meta":       "Meta",
	"nav":        "Nav",
	"ol":         "OL",
	"optgroup":   "Optgroup",
	"option":     "Option",
	"p":          "P",
	"section":    "Section",
	"select":     "Select",
	"small":      "Small",
	"span":       "Span",
	"strong":     "Strong",
	"summary":    "Summary",
	"table":      "Table",
	"tbody":      "TBody",
	"td":         "TD",
	"textarea":   "Textarea",
	"tfoot":      "TFoot",
	"th":         "TH",
	"thead":      "THead",
	"title":      "Title",
	"tr":         "TR",
	"ul":         "UL",
}

var boolAttrs = map[string]struct{}{
	"checked":    {},
	"defer":      {},
	"disabled":   {},
	"hidden":     {},
	"multiple":   {},
	"novalidate": {},
	"readonly":   {},
	"required":   {},
	"selected":   {},
}
