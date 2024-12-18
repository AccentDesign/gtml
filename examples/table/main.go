package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

var (
	tableCss   = "w-full caption-bottom text-sm"
	captionCss = "mt-4 text-sm text-gray-500"
	headCss    = "border-b"
	bodyCss    = "[&_tr:last-child]:border-0"
	footCss    = "border-t bg-gray-50 font-medium [&>tr]:last:border-b-0"
	trCss      = "border-b transition-colors hover:bg-gray-50"
	thCss      = "html-12 px-4 text-left align-middle font-medium"
	tdCss      = "p-4 align-middle font-medium"
)

var caption = gtml.Caption(gtml.Attrs{"class": captionCss}, gtml.Text("caption"))

var head = gtml.THead(
	gtml.Attrs{"class": headCss},
	gtml.TR(
		gtml.Attrs{"class": trCss},
		gtml.TH(gtml.Attrs{"class": thCss}, gtml.Text("col 1")),
		gtml.TH(gtml.Attrs{"class": thCss}, gtml.Text("col 2")),
		gtml.TH(gtml.Attrs{"class": thCss}, gtml.Text("col 3")),
		gtml.TH(gtml.Attrs{"class": thCss}, gtml.Text("col 4")),
	),
)

var body = gtml.TBody(
	gtml.Attrs{"class": bodyCss},
	gtml.TR(
		gtml.Attrs{"class": trCss},
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 1")),
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 2")),
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 3")),
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 4")),
	),
)

var foot = gtml.TFoot(
	gtml.Attrs{"class": footCss},
	gtml.TR(
		gtml.Attrs{"class": trCss},
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 1")),
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 2")),
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 3")),
		gtml.TD(gtml.Attrs{"class": tdCss}, gtml.Text("col 4")),
	),
)

var table = gtml.Table(gtml.Attrs{"class": tableCss}, caption, head, body, foot)

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = table.Render(context.Background(), os.Stdout)
}
