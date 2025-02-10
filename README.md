[![Test](https://github.com/AccentDesign/gtml/actions/workflows/test.yml/badge.svg)](https://github.com/AccentDesign/gtml/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/accentdesign/gtml)](https://goreportcard.com/report/github.com/accentdesign/gtml)
<a href="https://pkg.go.dev/github.com/accentdesign/gtml"><img src="https://img.shields.io/badge/Documentation%20on-pkg.go.dev-blue.svg"/></a>

# gtml

No-frills html elements in pure golang.

### Usage

```go
import "github.com/accentdesign/gtml"

field := gtml.Div(
    gtml.NA,
    gtml.Label(gtml.Attrs{"for": "email"}, gtml.Text("Email")),
    gtml.Input(gtml.Attrs{
        "class":       "form-input",
        "id":          "email",
        "placeholder": "email@example.com",
        "required":    true,
        "type":        "email",
    }),
    gtml.P(gtml.Attrs{"class": "help-text"}, gtml.Text("Your email address.")),
)

field.Render(context.Background(), os.Stdout)
```

renders:
```html
<div>
    <label for="email">Email</label>
    <input class="form-input" id="email" placeholder="email@example.com" required type="email">
    <p class="help-text">Your email address.</p>
</div>
```

### Templ

gtml is designed to work with [templ](https://templ.guide). It implements the `templ.Component` interface.

Easily include them in a `templ.Component`.

```go
templ Form() {
	<form>
		@field
	</form>
}
```

Or include a `templ.Component` in a `gtml.Element` as a child.

```go
func childTemplate(text string) templ.Component {
    return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
        _, err := io.WriteString(w, "<p>"+text+"</p>")
        return err
    })
}

gtml.Div(gtml.Attrs{"class": "parent"}, childTemplate("Hello, World!"))
```

### CLI

gtml also comes with a CLI tool to generate go from an html file.

install:
```sh
go install github.com/accentdesign/gtml/cmd/gtml@latest
```

usage:
```sh
gtml < input.html > output.go
```