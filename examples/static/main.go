package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"log"
	"os"
)

func hello(name string) *gtml.Element {
	return gtml.Div(gtml.NA, gtml.Text(fmt.Sprintf("Hello, %s!", name)))
}

func main() {
	f, err := os.Create("hello.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = hello("John").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
