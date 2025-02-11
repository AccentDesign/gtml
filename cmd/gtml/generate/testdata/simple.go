package html

import (
	. "github.com/accentdesign/gtml"
)

func Component() *Element {
	return Div(NA, Text("Hello world!"))
}
