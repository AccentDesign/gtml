package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

func Nav(isLoggedIn bool) *gtml.Element {
	return gtml.UL(
		gtml.NA,
		gtml.LI(gtml.NA, gtml.A(gtml.NA, gtml.Text("Home"))),
		gtml.If(isLoggedIn, gtml.LI(gtml.NA, gtml.A(gtml.NA, gtml.Text("Profile")))),
	)
}

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	nav := Nav(true)

	_ = nav.Render(context.Background(), os.Stdout)
}
