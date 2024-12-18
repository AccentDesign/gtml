package main

import (
	"context"
	"fmt"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

var (
	btnCls    = "inline-flex items-center justify-center rounded-md html-10 px-4 py-2 text-sm font-medium bg-black text-white hover:bg-black/80"
	fieldCls  = "space-y-2"
	footerCls = "flex items-center justify-between"
	formCls   = "grid gap-10"
	helpCls   = "text-sm text-gray-500"
	inputCls  = "flex w-full rounded-md border px-3 py-2 text-sm"
	labelCls  = "text-sm font-medium leading-none"
	linkCls   = "text-blue-600 hover:underline"
)

var input = gtml.Div(
	gtml.Attrs{"class": fieldCls},
	gtml.Label(gtml.Attrs{"class": labelCls, "for": "email"}, gtml.Text("Email")),
	gtml.Input(gtml.Attrs{
		"class":       inputCls,
		"id":          "email",
		"placeholder": "email@example.com",
		"required":    true,
		"type":        "email",
	}),
	gtml.P(gtml.Attrs{"class": helpCls}, gtml.Text("This is your email address.")),
)

var dropdown = gtml.Div(
	gtml.Attrs{"class": fieldCls},
	gtml.Label(gtml.Attrs{"class": labelCls, "for": "sex"}, gtml.Text("Sex")),
	gtml.Select(
		gtml.Attrs{"class": inputCls, "id": "sex"},
		gtml.Option(gtml.Attrs{"value": ""}, gtml.Text("Please Select...")),
		gtml.Option(gtml.Attrs{"value": "male"}, gtml.Text("Male")),
		gtml.Option(gtml.Attrs{"value": "female"}, gtml.Text("Female")),
	),
	gtml.P(gtml.Attrs{"class": helpCls}, gtml.Text("This is your sex.")),
)

var textarea = gtml.Div(
	gtml.Attrs{"class": fieldCls},
	gtml.Label(gtml.Attrs{"class": labelCls, "for": "comments"}, gtml.Text("Comments")),
	gtml.Textarea(gtml.Attrs{
		"class":       inputCls,
		"id":          "comments",
		"placeholder": "comments",
	}),
	gtml.P(gtml.Attrs{"class": helpCls}, gtml.Text("what do you want to comment?")),
)

var footer = gtml.Div(
	gtml.Attrs{"class": footerCls},
	gtml.Button(gtml.Attrs{"class": btnCls, "type": "submit"}, gtml.Text("Submit")),
	gtml.A(gtml.Attrs{"class": linkCls, "href": "#"}, gtml.Text("Some random link")),
)

var form = gtml.Form(gtml.Attrs{"class": formCls}, input, dropdown, textarea, footer)

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = form.Render(context.Background(), os.Stdout)
}
