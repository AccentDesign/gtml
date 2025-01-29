package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

type option struct {
	value, label string
}

func input(id, label, placeholder, typ, help string, required bool) *gtml.Element {
	return gtml.Div(
		gtml.Attrs{"class": "field"},
		gtml.Label(gtml.Attrs{"for": id}, gtml.Text(label)),
		gtml.Input(gtml.Attrs{"class": "input", "id": id, "placeholder": placeholder, "required": required, "type": typ}),
		gtml.If(help != "", gtml.P(gtml.Attrs{"class": "help"}, gtml.Text(help))),
	)
}

func textarea(id, label, placeholder, help string, required bool) *gtml.Element {
	return gtml.Div(
		gtml.Attrs{"class": "field"},
		gtml.Label(gtml.Attrs{"for": id}, gtml.Text(label)),
		gtml.Textarea(gtml.Attrs{"class": "input", "id": id, "placeholder": placeholder, "required": required}),
		gtml.If(help != "", gtml.P(gtml.Attrs{"class": "help"}, gtml.Text(help))),
	)
}

func dropdown(id, label, help string, required bool, options []option) *gtml.Element {
	return gtml.Div(
		gtml.Attrs{"class": "field"},
		gtml.Label(gtml.Attrs{"for": id}, gtml.Text(label)),
		gtml.Select(
			gtml.Attrs{"class": "input", "id": id, "required": required},
			gtml.Map[option](options, func(o option) templ.Component {
				return gtml.Option(gtml.Attrs{"value": o.value}, gtml.Text(o.label))
			})...,
		),
		gtml.If(help != "", gtml.P(gtml.Attrs{"class": "help"}, gtml.Text(help))),
	)
}

var genderOptions = []option{
	{value: "", label: "Please Select..."},
	{value: "male", label: "Male"},
	{value: "female", label: "Female"},
}

var form = gtml.Form(
	gtml.Attrs{"class": "form"},
	gtml.Fieldset(
		gtml.NA,
		gtml.Legend(gtml.NA, gtml.Text("Your Details")),
		input("name", "Name", "eg: Sam Smith", "text", "Enter your name.", true),
		input("email", "Email", "eg: sam@example.com", "email", "Enter your email.", true),
		dropdown("gender", "Gender", "Enter your gender.", false, genderOptions),
		textarea("comments", "Comments", "eg: Please sent me your number.", "What do you want?", false),
	),
	gtml.Div(
		gtml.Attrs{"class": "footer"},
		gtml.Button(gtml.Attrs{"class": "button", "type": "submit"}, gtml.Text("Submit")),
		gtml.A(gtml.Attrs{"href": "#"}, gtml.Text("Some random link")),
	),
)

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = form.Render(context.Background(), os.Stdout)
}
