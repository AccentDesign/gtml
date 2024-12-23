package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
	"os"
	"time"
)

func field(controls ...templ.Component) *gtml.Element {
	return gtml.Div(gtml.Attrs{"class": "space-y-2"}, controls...)
}

func input(id, label, placeholder, typ, help string, required bool) *gtml.Element {
	return field(
		gtml.Label(gtml.Attrs{"class": "text-sm font-medium leading-none", "for": id}, gtml.Text(label)),
		gtml.Input(gtml.Attrs{"class": "flex w-full rounded-md border px-3 py-2 text-sm", "id": id, "placeholder": placeholder, "required": required, "type": typ}),
		gtml.P(gtml.Attrs{"class": "text-sm text-gray-500"}, gtml.Text(help)),
	)
}

func textarea(id, label, placeholder, help string, required bool) *gtml.Element {
	return field(
		gtml.Label(gtml.Attrs{"class": "text-sm font-medium leading-none", "for": id}, gtml.Text(label)),
		gtml.Textarea(gtml.Attrs{"class": "flex w-full rounded-md border px-3 py-2 text-sm", "id": id, "placeholder": placeholder, "required": required}),
		gtml.P(gtml.Attrs{"class": "text-sm text-gray-500"}, gtml.Text(help)),
	)
}

func dropdown(id, label, help string, required bool, options []templ.Component) *gtml.Element {
	return field(
		gtml.Label(gtml.Attrs{"class": "text-sm font-medium leading-none", "for": id}, gtml.Text(label)),
		gtml.Select(gtml.Attrs{"class": "flex w-full rounded-md border px-3 py-2 text-sm", "id": id, "required": required}, options...),
		gtml.P(gtml.Attrs{"class": "text-sm text-gray-500"}, gtml.Text(help)),
	)
}

var sexOptions = []templ.Component{
	gtml.Option(gtml.Attrs{"value": ""}, gtml.Text("Please Select...")),
	gtml.Option(gtml.Attrs{"value": "male"}, gtml.Text("Male")),
	gtml.Option(gtml.Attrs{"value": "female"}, gtml.Text("Female")),
	gtml.Option(gtml.Attrs{"value": "other"}, gtml.Text("Other")),
}

var form = gtml.Form(
	gtml.Attrs{"class": "grid gap-10"},
	gtml.Fieldset(
		gtml.Attrs{"class": "grid gap-10"},
		gtml.Legend(gtml.Attrs{"class": "mb-10"}, gtml.Text("Your Details")),
		input("name", "Name", "eg: Sam Smith", "text", "Enter your full name.", true),
		input("email", "Email", "eg: sam@example.com", "email", "Enter your email address.", true),
		dropdown("sex", "Sex", "This is your sex.", false, sexOptions),
		textarea("comments", "Comments", "eg: Please sent me your number.", "What do you want?", false),
	),
	gtml.Div(
		gtml.Attrs{"class": "flex items-center justify-between"},
		gtml.Button(gtml.Attrs{"class": "rounded-md px-4 py-2 text-sm font-medium bg-black text-white", "type": "submit"}, gtml.Text("Submit")),
		gtml.A(gtml.Attrs{"class": "text-blue-600 hover:underline", "href": "#"}, gtml.Text("Some random link")),
	),
)

func main() {
	defer func(start time.Time) {
		fmt.Println("")
		fmt.Println(time.Since(start))
	}(time.Now())

	_ = form.Render(context.Background(), os.Stdout)
}
