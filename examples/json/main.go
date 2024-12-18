package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

var ul = gtml.UL(
	gtml.NA,
	gtml.LI(gtml.NA, gtml.Text("Hello")),
	gtml.LI(gtml.NA, gtml.Text("World")),
)

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	// to json
	jsonData, err := json.MarshalIndent(ul, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	// from json
	var element gtml.Element
	if err := json.Unmarshal(jsonData, &element); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// render
	if err := element.Render(context.Background(), os.Stdout); err != nil {
		fmt.Println("Error rendering HTML:", err)
	}
}
