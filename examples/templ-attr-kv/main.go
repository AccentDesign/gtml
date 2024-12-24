package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

var body = gtml.Body(gtml.Attrs{"class": templ.KV("dark", true)})

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = body.Render(context.Background(), os.Stdout)
}
