package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"io"
	"os"
	"time"
)

type Table struct {
	CurrentPage int
	Limit       int
	TotalPages  int
}

func (t *Table) Header() *gtml.Element {
	return gtml.THead(
		gtml.NA,
		gtml.TR(
			gtml.NA,
			gtml.TH(gtml.NA, gtml.Text("Number")),
		),
	)
}

func (t *Table) Body() *gtml.Element {
	body := gtml.TBody(gtml.NA)
	start := (t.CurrentPage-1)*t.Limit + 1
	end := t.CurrentPage * t.Limit
	for i := start; i <= end; i++ {
		body.AddChildren(gtml.TR(
			gtml.NA,
			gtml.TD(gtml.NA, gtml.Text(fmt.Sprintf("%d", i))),
		))
	}
	return body
}

func (t *Table) Footer() *gtml.Element {
	if t.CurrentPage == t.TotalPages {
		return &gtml.Element{}
	}
	return gtml.TFoot(
		gtml.NA,
		gtml.TD(
			gtml.NA,
			gtml.Button(
				gtml.NA,
				gtml.Text(fmt.Sprintf("Page %d", t.CurrentPage+1)),
			),
		),
	)
}

func (t *Table) Render(ctx context.Context, w io.Writer) error {
	tbl := gtml.Table(
		gtml.NA,
		t.Header(),
		t.Body(),
		t.Footer(),
	)
	return tbl.Render(ctx, w)
}

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	table := Table{
		CurrentPage: 1,
		Limit:       10,
		TotalPages:  10,
	}

	for page := 1; page <= table.TotalPages; page++ {
		table.CurrentPage = page
		fmt.Println("")
		_ = table.Render(context.Background(), os.Stdout)
	}
}
