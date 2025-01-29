package gtml

import (
	"github.com/a-h/templ"
)

// If returns the element if the condition is true, otherwise returns an empty element.
func If(condition bool, element *Element) *Element {
	if condition {
		return element
	}
	return Empty()
}

// IfElse returns the truthy element if the condition is true, otherwise returns the falsy element.
func IfElse(condition bool, truthyElement, falsyElement *Element) *Element {
	if condition {
		return truthyElement
	}
	return falsyElement
}

// Map function that applies a transformation function to a list of items and returns the resulting elements.
func Map[T any](items []T, transform func(item T) templ.Component) []templ.Component {
	elements := make([]templ.Component, len(items))
	for i, item := range items {
		elements[i] = transform(item)
	}
	return elements
}

// Range function that generates a sequence of numbers and renders elements based on those numbers.
func Range(start, end int, transform func(i int) templ.Component) []templ.Component {
	elements := make([]templ.Component, end-start)
	for i := start; i < end; i++ {
		elements[i-start] = transform(i)
	}
	return elements
}
