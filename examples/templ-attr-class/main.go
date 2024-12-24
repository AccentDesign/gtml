package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

var body = gtml.Body(gtml.Attrs{"class": class("unique", "css", "css")})

// class returns a string of classes, de-duped by templ.
func class(classes ...string) string {
	return templ.Classes(classes).String()
}

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = body.Render(context.Background(), os.Stdout)
}
