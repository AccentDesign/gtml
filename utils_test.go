package gtml

import (
	"fmt"
	"github.com/a-h/templ"
	"testing"
)

// Test the If function with true and false conditions.
func TestIfFunction(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		element   *Element
		expected  string
	}{
		{"IfTrue", true, Div(NA, Text("Visible")), `<div>Visible</div>`},
		{"IfFalse", false, Div(NA, Text("Visible")), ``},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := If(tt.condition, tt.element)
			output := renderElement(t, node)
			if output != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, output)
			}
		})
	}
}

// Test the IfElse function with true and false conditions.
func TestIfElseFunction(t *testing.T) {
	tests := []struct {
		name          string
		condition     bool
		truthyElement *Element
		falsyElement  *Element
		expected      string
	}{
		{"IfElseTrue", true, Div(NA, Text("Visible")), Div(NA, Text("Hidden")), `<div>Visible</div>`},
		{"IfElseFalse", false, Div(NA, Text("Visible")), Div(NA, Text("Hidden")), `<div>Hidden</div>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := IfElse(tt.condition, tt.truthyElement, tt.falsyElement)
			output := renderElement(t, node)
			if output != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, output)
			}
		})
	}
}

// Test the Map function with a list of items.
func TestMapFunction(t *testing.T) {
	items := []string{"One", "Two", "Three"}
	transform := func(item string) templ.Component {
		return LI(NA, Text(item))
	}

	node := UL(NA, Map[string](items, transform)...)
	output := renderElement(t, node)
	expected := `<ul><li>One</li><li>Two</li><li>Three</li></ul>`

	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// Test the Range function with a start and end value.
func TestRangeFunction(t *testing.T) {
	transform := func(i int) templ.Component {
		return LI(NA, Text(fmt.Sprintf("%d", i)))
	}

	node := UL(NA, Range(1, 4, transform)...)
	output := renderElement(t, node)
	expected := `<ul><li>1</li><li>2</li><li>3</li></ul>`

	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
