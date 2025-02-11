package html

import (
	. "github.com/accentdesign/gtml"
)

func Component() *Element {
	return New("custom", Attrs{"class": "custom"}, Text("Tag"))
}
