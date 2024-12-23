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
    <input class="form-input" id="email" placeholder="email@example.com" required type="email" />
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