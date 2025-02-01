package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"github.com/accentdesign/gtml/attrs"
	"os"
	"time"
)

var field = gtml.Div(
	attrs.Class("field"),
	gtml.Label(attrs.For("id_field"), gtml.Text("Field")),
	gtml.Input(attrs.Merge(
		attrs.Class("input"),
		attrs.ID("id_field"),
		attrs.Placeholder("Enter something..."),
		attrs.Required(true),
		attrs.Type("text"),
	)),
	gtml.P(attrs.Class("help"), gtml.Text("Help text.")),
)

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = field.Render(context.Background(), os.Stdout)
}
