package gtml

import "testing"

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
