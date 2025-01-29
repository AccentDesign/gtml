package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

var caption = gtml.Caption(gtml.NA, gtml.Text("caption"))

var head = gtml.THead(
	gtml.NA,
	gtml.TR(
		gtml.NA,
		gtml.TH(gtml.NA, gtml.Text("col 1")),
		gtml.TH(gtml.NA, gtml.Text("col 2")),
		gtml.TH(gtml.NA, gtml.Text("col 3")),
		gtml.TH(gtml.NA, gtml.Text("col 4")),
	),
)

var body = gtml.TBody(
	gtml.NA,
	gtml.TR(
		gtml.NA,
		gtml.TD(gtml.NA, gtml.Text("col 1")),
		gtml.TD(gtml.NA, gtml.Text("col 2")),
		gtml.TD(gtml.NA, gtml.Text("col 3")),
		gtml.TD(gtml.NA, gtml.Text("col 4")),
	),
)

var foot = gtml.TFoot(
	gtml.NA,
	gtml.TR(
		gtml.NA,
		gtml.TD(gtml.NA, gtml.Text("col 1")),
		gtml.TD(gtml.NA, gtml.Text("col 2")),
		gtml.TD(gtml.NA, gtml.Text("col 3")),
		gtml.TD(gtml.NA, gtml.Text("col 4")),
	),
)

var table = gtml.Table(gtml.NA, caption, head, body, foot)

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = table.Render(context.Background(), os.Stdout)
}
