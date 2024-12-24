package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
	"io"
	"os"
	"time"
)

func childTemplate(text string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, "<p>"+text+"</p>")
		return err
	})
}

var div = gtml.Div(gtml.Attrs{"class": "parent"}, childTemplate("Hello, World!"))

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = div.Render(context.Background(), os.Stdout)
}
