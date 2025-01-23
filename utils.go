package gtml

// If returns the element if the condition is true, otherwise returns an empty element.
func If(condition bool, element *Element) *Element {
	if condition {
		return element
	}
	return Empty()
}
