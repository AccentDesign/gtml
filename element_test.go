package gtml

import (
	"bytes"
	"context"
	"testing"
)

// Helper function to simplify rendering a Element and capturing the output.
func renderElement(t *testing.T, node *Element) string {
	var buf bytes.Buffer
	err := node.Render(context.Background(), &buf)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}
	return buf.String()
}

// Test to ensure each helper function uses the correct tag.
func TestElementTags(t *testing.T) {
	tests := []struct {
		name     string
		nodeFunc func() *Element
		expected string
	}{
		{"A", func() *Element { return A(NA) }, "a"},
		{"Aside", func() *Element { return Aside(NA) }, "aside"},
		{"Button", func() *Element { return Button(NA) }, "button"},
		{"BR", func() *Element { return BR(NA) }, "br"},
		{"Caption", func() *Element { return Caption(NA) }, "caption"},
		{"Details", func() *Element { return Details(NA) }, "details"},
		{"Div", func() *Element { return Div(NA) }, "div"},
		{"Form", func() *Element { return Form(NA) }, "form"},
		{"H1", func() *Element { return H1(NA) }, "h1"},
		{"H2", func() *Element { return H2(NA) }, "h2"},
		{"H3", func() *Element { return H3(NA) }, "h3"},
		{"H4", func() *Element { return H4(NA) }, "h4"},
		{"H5", func() *Element { return H5(NA) }, "h5"},
		{"Header", func() *Element { return Header(NA) }, "header"},
		{"HR", func() *Element { return HR(NA) }, "hr"},
		{"Img", func() *Element { return Img(NA) }, "img"},
		{"Input", func() *Element { return Input(NA) }, "input"},
		{"Label", func() *Element { return Label(NA) }, "label"},
		{"LI", func() *Element { return LI(NA) }, "li"},
		{"Main", func() *Element { return Main(NA) }, "main"},
		{"Nav", func() *Element { return Nav(NA) }, "nav"},
		{"OL", func() *Element { return OL(NA) }, "ol"},
		{"Option", func() *Element { return Option(NA) }, "option"},
		{"P", func() *Element { return P(NA) }, "p"},
		{"Section", func() *Element { return Section(NA) }, "section"},
		{"Select", func() *Element { return Select(NA) }, "select"},
		{"Span", func() *Element { return Span(NA) }, "span"},
		{"Summary", func() *Element { return Summary(NA) }, "summary"},
		{"Table", func() *Element { return Table(NA) }, "table"},
		{"TBody", func() *Element { return TBody(NA) }, "tbody"},
		{"TD", func() *Element { return TD(NA) }, "td"},
		{"Textarea", func() *Element { return Textarea(NA) }, "textarea"},
		{"TFoot", func() *Element { return TFoot(NA) }, "tfoot"},
		{"TH", func() *Element { return TH(NA) }, "th"},
		{"THead", func() *Element { return THead(NA) }, "thead"},
		{"TR", func() *Element { return TR(NA) }, "tr"},
		{"UL", func() *Element { return UL(NA) }, "ul"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := tt.nodeFunc()
			if node.Tag != tt.expected {
				t.Errorf("Expected tag %q, got %q", tt.expected, node.Tag)
			}
		})
	}
}

// Test rendering a self-closing Element.
func TestRenderSelfClosingElements(t *testing.T) {
	tests := []struct {
		name     string
		nodeFunc func() *Element
		expected string
	}{
		{"Br", func() *Element { return BR(NA) }, "<br />"},
		{"Img", func() *Element { return Img(NA) }, "<img />"},
		{"Input", func() *Element { return Input(NA) }, "<input />"},
		{"Hr", func() *Element { return HR(NA) }, "<hr />"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := tt.nodeFunc()
			output := renderElement(t, node)
			if output != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, output)
			}
		})
	}
}

// Test rendering a basic Element with a tag and text content.
func TestRenderBasicElement(t *testing.T) {
	node := Div(Attrs{"class": "example"}, Text("Hello, World!"))
	expected := `<div class="example">Hello, World!</div>`

	output := renderElement(t, node)
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test rendering an image tag with attributes.
func TestRenderImgElement(t *testing.T) {
	node := Img(Attrs{"src": "image.png", "alt": "Test Image"})
	expected := `<img alt="Test Image" src="image.png" />`

	output := renderElement(t, node)
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test rendering nested Element's with children.
func TestRenderNestedElements(t *testing.T) {
	node := Div(Attrs{"class": "parent"},
		Span(Attrs{"class": "child"}, Text("Child Element")),
	)
	expected := `<div class="parent"><span class="child">Child Element</span></div>`

	output := renderElement(t, node)
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test rendering an Element with text content but no tag.
func TestRenderTextElement(t *testing.T) {
	node := Text("Just text")
	expected := `Just text`

	output := renderElement(t, node)
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test rendering an Element with multiple attributes.
func TestRenderElementWithAttributes(t *testing.T) {
	node := Div(Attrs{"id": "main", "data-test": "123"}, Text("Content"))
	expected := `<div data-test="123" id="main">Content</div>`

	output := renderElement(t, node)
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test rendering empty Element (no tag, no attributes, no children).
func TestRenderEmptyElement(t *testing.T) {
	node := &Element{}
	expected := ``

	output := renderElement(t, node)
	if output != expected {
		t.Errorf("Expected an empty string, got %q", output)
	}
}

// Test adding children to an Element
func TestRenderElementAddChildren(t *testing.T) {
	node := UL(NA, LI(NA, Text("1")))
	node.AddChildren(LI(NA, Text("2")), LI(NA, Text("3")))
	expected := `<ul><li>1</li><li>2</li><li>3</li></ul>`

	output := renderElement(t, node)
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
