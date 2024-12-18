# gtml

No-frills html from golang.

Example:

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

With [templ](https://templ.guide).

```go
templ Form() {
	<form>
		@field
	</form>
}
```