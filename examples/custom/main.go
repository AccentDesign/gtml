package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

func icon(children ...*gtml.Element) *gtml.Element {
	return &gtml.Element{
		Tag: "svg",
		Attrs: gtml.Attrs{
			"width":   "32",
			"height":  "32",
			"xmlns":   "http://www.w3.org/2000/svg",
			"viewBox": "0 0 24 24",
			"fill":    "currentColor",
		},
		Children: children,
	}
}

var check = icon(&gtml.Element{
	Tag: "path",
	Attrs: gtml.Attrs{
		"fill":            "none",
		"stroke":          "currentColor",
		"stroke-linecap":  "round",
		"stroke-linejoin": "round",
		"stroke-width":    "2",
		"d":               "M20 6L9 17l-5-5",
	},
})

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = check.Render(context.Background(), os.Stdout)
}
