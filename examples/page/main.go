package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

func page(title string, content ...templ.Component) *gtml.Element {
	return &gtml.Element{
		Children: []templ.Component{
			gtml.Raw("<!DOCTYPE html>"),
			gtml.Html(
				gtml.Attrs{"lang": "en"},
				gtml.Head(
					gtml.NA,
					gtml.Meta(gtml.Attrs{"charset": "UTF-8"}),
					gtml.Title(gtml.NA, gtml.Text(title)),
					gtml.Link(gtml.Attrs{"rel": "stylesheet", "href": "style.css"}),
				),
				gtml.Body(gtml.NA, content...),
			),
		},
	}
}

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	heading := gtml.H1(gtml.NA, gtml.Text("Hello World!"))
	_ = page("Hello World!", heading).Render(context.Background(), os.Stdout)
}
