package attrs

import (
	"bytes"
	"context"
	"github.com/accentdesign/gtml"
	"testing"
)

func TestStringAttributes(t *testing.T) {
	tests := []struct {
		name          string
		fnc           func(string) gtml.Attrs
		expectedAttr  string
		expectedValue string
	}{
		{"Accept", Accept, "accept", "image/*"},
		{"AccessKey", AccessKey, "accesskey", "a"},
		{"Action", Action, "action", "a"},
		{"Alt", Alt, "alt", "a"},
		{"AutoComplete", AutoComplete, "autocomplete", "on"},
		{"Charset", Charset, "charset", "utf-8"},
		{"ColSpan", ColSpan, "colspan", "2"},
		{"Data", Data, "data", "d"},
		{"Dir", Dir, "dir", "ltr"},
		{"EncType", EncType, "enctype", "multipart/form-data"},
		{"For", For, "for", "i"},
		{"Form", Form, "form", "f"},
		{"Height", Height, "height", "100"},
		{"Href", Href, "href", "/home"},
		{"HxBoost", HxBoost, "hx-boost", "true"},
		{"HxConfirm", HxConfirm, "hx-confirm", "Are you sure?"},
		{"HxDelete", HxDelete, "hx-delete", "/home"},
		{"HxGet", HxGet, "hx-get", "/home"},
		{"HxInclude", HxInclude, "hx-include", "#id"},
		{"HxInherit", HxInherit, "hx-inherit", "hx-target"},
		{"HxPost", HxPost, "hx-post", "/home"},
		{"HxPushUrl", HxPushUrl, "hx-push-url", "true"},
		{"HxReplaceUrl", HxReplaceUrl, "hx-replace-url", "true"},
		{"HxSelect", HxSelect, "hx-select", "#id"},
		{"HxSelectOob", HxSelectOob, "hx-select-oob", "#id"},
		{"HxSwap", HxSwap, "hx-swap", "outerHTML"},
		{"HxSwapOob", HxSwapOob, "hx-swap-oob", "true"},
		{"hxTarget", HxTarget, "hx-target", "#id"},
		{"HxTrigger", HxTrigger, "hx-trigger", "click"},
		{"ID", ID, "id", "i"},
		{"Lang", Lang, "lang", "en"},
		{"Max", Max, "max", "10"},
		{"MaxLength", MaxLength, "maxlength", "10"},
		{"Media", Media, "media", "screen"},
		{"Method", Method, "method", "get"},
		{"Min", Min, "min", "10"},
		{"MinLength", MinLength, "minlength", "10"},
		{"Name", Name, "name", "foo"},
		{"Pattern", Pattern, "pattern", "w{3,16}"},
		{"Placeholder", Placeholder, "placeholder", "foo"},
		{"Rel", Rel, "rel", "stylesheet"},
		{"Role", Role, "role", "button"},
		{"RowSpan", RowSpan, "rowspan", "2"},
		{"Size", Size, "size", "10"},
		{"Sizes", Sizes, "sizes", "(max-width: 600px) 480px, 800px"},
		{"Span", Span, "span", "2"},
		{"SpellCheck", SpellCheck, "spellcheck", "true"},
		{"Src", Src, "src", "image.png"},
		{"SrcSet", SrcSet, "srcset", "elva-fairy-480w.jpg 480w, elva-fairy-800w.jpg 800w"},
		{"Step", Step, "step", "10"},
		{"Style", Style, "style", "width:100%"},
		{"TabIndex", TabIndex, "tabindex", "1"},
		{"Target", Target, "target", "_blank"},
		{"Title", Title, "title", "t"},
		{"Translate", Translate, "translate", "yes"},
		{"Type", Type, "type", "text"},
		{"Name", Name, "name", "foo"},
		{"Value", Value, "value", "foo"},
		{"Width", Width, "width", "100"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			attr := test.fnc(test.expectedValue)
			if attr[test.expectedAttr] != test.expectedValue {
				t.Errorf("Expected %s to be %s, got %s", test.expectedAttr, test.expectedValue, attr[test.expectedAttr])
			}
		})
	}
}

func TestVariadicAttributes(t *testing.T) {
	tests := []struct {
		name          string
		fnc           func(...string) gtml.Attrs
		params        []string
		expectedAttr  string
		expectedValue string
	}{
		{"Class", Class, []string{"foo", "bar"}, "class", "foo bar"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			attr := test.fnc(test.params...)
			if attr[test.expectedAttr] != test.expectedValue {
				t.Errorf("Expected %s to be %s, got %s", test.expectedAttr, test.expectedValue, attr[test.expectedAttr])
			}
		})
	}
}

func TestBooleanAttributes(t *testing.T) {
	tests := []struct {
		name          string
		fnc           func(bool) gtml.Attrs
		expectedAttr  string
		expectedValue bool
	}{
		{"AutoFocus", AutoFocus, "autofocus", true},
		{"Checked", Checked, "checked", true},
		{"Disabled", Disabled, "disabled", true},
		{"Hidden", Hidden, "hidden", true},
		{"Multiple", Multiple, "multiple", true},
		{"NoValidate", NoValidate, "novalidate", true},
		{"ReadOnly", ReadOnly, "readonly", true},
		{"Required", Required, "required", true},
		{"Selected", Selected, "selected", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			attr := test.fnc(test.expectedValue)
			if attr[test.expectedAttr] != test.expectedValue {
				t.Errorf("Expected %s to be %v, got %v", test.expectedAttr, test.expectedValue, attr[test.expectedAttr])
			}
		})
	}
}

func TestAriaAttributes(t *testing.T) {
	t.Run("Aria prefixes attributes with 'aria-'", func(t *testing.T) {
		res := Aria("selected", "true")
		if res["aria-selected"] != "true" {
			t.Errorf("Expected aria-selected to be 'true', got %s", res["aria-selected"])
		}
	})
}

func TestEmpty(t *testing.T) {
	a := Empty()
	if len(a) != 0 {
		t.Errorf("Expected no attributes, got %v", a)
	}
}

func TestIf(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		expected  gtml.Attrs
	}{
		{"true", true, gtml.Attrs{"foo": "bar"}},
		{"false", false, gtml.Attrs{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a := If(test.condition, gtml.Attrs{"foo": "bar"})
			if len(a) != len(test.expected) {
				t.Errorf("Expected %d attributes, got %d", len(test.expected), len(a))
			}
			for k, v := range test.expected {
				if a[k] != v {
					t.Errorf("Expected %s to be %s, got %s", k, v, a[k])
				}
			}
		})
	}
}

func TestMerge(t *testing.T) {
	a := Merge(ID("some-id"), Class("foo", "bar"))
	if a["id"] != "some-id" {
		t.Errorf("Expected id to be 'some-id', got %s", a["id"])
	}
	if a["class"] != "foo bar" {
		t.Errorf("Expected class to be 'foo bar', got %s", a["class"])
	}
}

func BenchmarkBasicAttrRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		node := gtml.Div(gtml.Attrs{"id": "div", "disabled": true}, gtml.Text("Hello"))
		var buf bytes.Buffer
		err := node.Render(context.Background(), &buf)
		if err != nil {
			b.Fatalf("Render failed: %v", err)
		}
	}
}

func BenchmarkFuncAttrRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		node := gtml.Div(Merge(ID("id"), Disabled(true)), gtml.Text("Hello"))
		var buf bytes.Buffer
		err := node.Render(context.Background(), &buf)
		if err != nil {
			b.Fatalf("Render failed: %v", err)
		}
	}
}
