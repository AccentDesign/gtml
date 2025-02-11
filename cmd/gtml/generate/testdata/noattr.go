package html

import (
	. "github.com/accentdesign/gtml"
)

func Component() *Element {
	return UL(NA, LI(NA, Text("Item 1")), LI(NA, Text("Item 2")))
}
