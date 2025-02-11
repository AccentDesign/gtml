package html

import (
	. "github.com/accentdesign/gtml"
)

func Component() *Element {
	return Form(Attrs{"novalidate": true}, Div(NA, Label(Attrs{"for": "email"}, Text("Email")), Input(Attrs{"class": "form-input", "id": "email", "placeholder": "email@example.com", "required": true, "type": "email"}), P(Attrs{"class": "help-text"}, Text("Your email address."))), Button(Attrs{"class": "btn", "type": "submit"}, Text("Submit")))
}
