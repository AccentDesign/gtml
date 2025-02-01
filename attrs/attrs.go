package attrs

import (
	"github.com/a-h/templ"
	"github.com/accentdesign/gtml"
)

func Accept(value string) gtml.Attrs {
	return gtml.Attrs{"accept": value}
}

func AccessKey(value string) gtml.Attrs {
	return gtml.Attrs{"accesskey": value}
}

func Action(value string) gtml.Attrs {
	return gtml.Attrs{"action": value}
}

func Alt(value string) gtml.Attrs {
	return gtml.Attrs{"alt": value}
}

func Aria(attr, value string) gtml.Attrs {
	return gtml.Attrs{"aria-" + attr: value}
}

func AutoComplete(value string) gtml.Attrs {
	return gtml.Attrs{"autocomplete": value}
}

func AutoFocus(value bool) gtml.Attrs {
	return gtml.Attrs{"autofocus": value}
}

func Charset(value string) gtml.Attrs {
	return gtml.Attrs{"charset": value}
}

func Checked(value bool) gtml.Attrs {
	return gtml.Attrs{"checked": value}
}

func Class(values ...string) gtml.Attrs {
	return gtml.Attrs{"class": templ.Classes(values).String()}
}

func ColSpan(value string) gtml.Attrs {
	return gtml.Attrs{"colspan": value}
}

func Data(value string) gtml.Attrs {
	return gtml.Attrs{"data": value}
}

func Dir(value string) gtml.Attrs {
	return gtml.Attrs{"dir": value}
}

func Disabled(value bool) gtml.Attrs {
	return gtml.Attrs{"disabled": value}
}

func EncType(value string) gtml.Attrs {
	return gtml.Attrs{"enctype": value}
}

func For(value string) gtml.Attrs {
	return gtml.Attrs{"for": value}
}

func Form(value string) gtml.Attrs {
	return gtml.Attrs{"form": value}
}

func Height(value string) gtml.Attrs {
	return gtml.Attrs{"height": value}
}

func Hidden(value bool) gtml.Attrs {
	return gtml.Attrs{"hidden": value}
}

func Href(value string) gtml.Attrs {
	return gtml.Attrs{"href": value}
}

func HxBoost(value string) gtml.Attrs {
	return gtml.Attrs{"hx-boost": value}
}

func HxConfirm(value string) gtml.Attrs {
	return gtml.Attrs{"hx-confirm": value}
}

func HxDelete(value string) gtml.Attrs {
	return gtml.Attrs{"hx-delete": value}
}

func HxGet(value string) gtml.Attrs {
	return gtml.Attrs{"hx-get": value}
}

func HxInclude(value string) gtml.Attrs {
	return gtml.Attrs{"hx-include": value}
}

func HxInherit(value string) gtml.Attrs {
	return gtml.Attrs{"hx-inherit": value}
}

func HxPost(value string) gtml.Attrs {
	return gtml.Attrs{"hx-post": value}
}

func HxPushUrl(value string) gtml.Attrs {
	return gtml.Attrs{"hx-push-url": value}
}

func HxReplaceUrl(value string) gtml.Attrs {
	return gtml.Attrs{"hx-replace-url": value}
}

func HxSelect(value string) gtml.Attrs {
	return gtml.Attrs{"hx-select": value}
}

func HxSelectOob(value string) gtml.Attrs {
	return gtml.Attrs{"hx-select-oob": value}
}

func HxSwap(value string) gtml.Attrs {
	return gtml.Attrs{"hx-swap": value}
}

func HxSwapOob(value string) gtml.Attrs {
	return gtml.Attrs{"hx-swap-oob": value}
}

func HxTarget(value string) gtml.Attrs {
	return gtml.Attrs{"hx-target": value}
}

func HxTrigger(value string) gtml.Attrs {
	return gtml.Attrs{"hx-trigger": value}
}

func ID(value string) gtml.Attrs {
	return gtml.Attrs{"id": value}
}

func Lang(value string) gtml.Attrs {
	return gtml.Attrs{"lang": value}
}

func Max(value string) gtml.Attrs {
	return gtml.Attrs{"max": value}
}

func MaxLength(value string) gtml.Attrs {
	return gtml.Attrs{"maxlength": value}
}

func Media(value string) gtml.Attrs {
	return gtml.Attrs{"media": value}
}

func Method(value string) gtml.Attrs {
	return gtml.Attrs{"method": value}
}

func Min(value string) gtml.Attrs {
	return gtml.Attrs{"min": value}
}

func MinLength(value string) gtml.Attrs {
	return gtml.Attrs{"minlength": value}
}

func Multiple(value bool) gtml.Attrs {
	return gtml.Attrs{"multiple": value}
}

func Name(value string) gtml.Attrs {
	return gtml.Attrs{"name": value}
}

func NoValidate(value bool) gtml.Attrs {
	return gtml.Attrs{"novalidate": value}
}

func Pattern(value string) gtml.Attrs {
	return gtml.Attrs{"pattern": value}
}

func Placeholder(value string) gtml.Attrs {
	return gtml.Attrs{"placeholder": value}
}

func ReadOnly(value bool) gtml.Attrs {
	return gtml.Attrs{"readonly": value}
}

func Rel(value string) gtml.Attrs {
	return gtml.Attrs{"rel": value}
}

func Required(value bool) gtml.Attrs {
	return gtml.Attrs{"required": value}
}

func Role(value string) gtml.Attrs {
	return gtml.Attrs{"role": value}
}

func RowSpan(value string) gtml.Attrs {
	return gtml.Attrs{"rowspan": value}
}

func Selected(value bool) gtml.Attrs {
	return gtml.Attrs{"selected": value}
}

func Size(value string) gtml.Attrs {
	return gtml.Attrs{"size": value}
}

func Sizes(value string) gtml.Attrs {
	return gtml.Attrs{"sizes": value}
}

func Span(value string) gtml.Attrs {
	return gtml.Attrs{"span": value}
}

func SpellCheck(value string) gtml.Attrs {
	return gtml.Attrs{"spellcheck": value}
}

func Src(value string) gtml.Attrs {
	return gtml.Attrs{"src": value}
}

func SrcSet(value string) gtml.Attrs {
	return gtml.Attrs{"srcset": value}
}

func Step(value string) gtml.Attrs {
	return gtml.Attrs{"step": value}
}

func Style(value string) gtml.Attrs {
	return gtml.Attrs{"style": value}
}

func TabIndex(value string) gtml.Attrs {
	return gtml.Attrs{"tabindex": value}
}

func Target(value string) gtml.Attrs {
	return gtml.Attrs{"target": value}
}

func Title(value string) gtml.Attrs {
	return gtml.Attrs{"title": value}
}

func Translate(value string) gtml.Attrs {
	return gtml.Attrs{"translate": value}
}

func Type(value string) gtml.Attrs {
	return gtml.Attrs{"type": value}
}

func Value(value string) gtml.Attrs {
	return gtml.Attrs{"value": value}
}

func Width(value string) gtml.Attrs {
	return gtml.Attrs{"width": value}
}

func Empty() gtml.Attrs {
	return gtml.Attrs{}
}

func If(cond bool, attrs gtml.Attrs) gtml.Attrs {
	if cond {
		return attrs
	}
	return Empty()
}

func Merge(attrs ...gtml.Attrs) gtml.Attrs {
	var res = gtml.Attrs{}
	for _, attr := range attrs {
		for k, v := range attr {
			res[k] = v
		}
	}
	return res
}
