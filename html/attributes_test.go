package html

import (
	"bytes"
	"fmt"
	"testing"

	g "github.com/fuguohong1024/gomponents"
	"github.com/fuguohong1024/gomponents/internal/assert"
)

func TestBooleanAttributes(t *testing.T) {
	tests := []struct {
		Name string
		Func func() g.Node
	}{
		{Name: "async", Func: Async},
		{Name: "autofocus", Func: AutoFocus},
		{Name: "autoplay", Func: AutoPlay},
		{Name: "checked", Func: Checked},
		{Name: "controls", Func: Controls},
		{Name: "defer", Func: Defer},
		{Name: "disabled", Func: Disabled},
		{Name: "loop", Func: Loop},
		{Name: "multiple", Func: Multiple},
		{Name: "muted", Func: Muted},
		{Name: "playsinline", Func: PlaysInline},
		{Name: "readonly", Func: ReadOnly},
		{Name: "required", Func: Required},
		{Name: "selected", Func: Selected},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			n := g.El("div", test.Func())
			b := new(bytes.Buffer)
			n.Render(b)
			assert.Equal(t, fmt.Sprintf(`<div %v></div>`, test.Name), b.String())
		})
	}
}

func TestSimpleAttributes(t *testing.T) {
	tests := []struct {
		Name string
		Func func(string) g.Node
	}{
		{Name: "accept", Func: Accept},
		{Name: "action", Func: Action},
		{Name: "alt", Func: Alt},
		{Name: "as", Func: As},
		{Name: "autocomplete", Func: AutoComplete},
		{Name: "charset", Func: Charset},
		{Name: "cite", Func: CiteAttr},
		{Name: "class", Func: Class},
		{Name: "cols", Func: Cols},
		{Name: "colspan", Func: ColSpan},
		{Name: "content", Func: Content},
		{Name: "crossorigin", Func: CrossOrigin},
		{Name: "datetime", Func: DateTime},
		{Name: "draggable", Func: Draggable},
		{Name: "enctype", Func: EncType},
		{Name: "dir", Func: Dir},
		{Name: "for", Func: For},
		{Name: "form", Func: FormAttr},
		{Name: "height", Func: Height},
		{Name: "hidden", Func: Hidden},
		{Name: "href", Func: Href},
		{Name: "id", Func: ID},
		{Name: "integrity", Func: Integrity},
		{Name: "label", Func: LabelAttr},
		{Name: "lang", Func: Lang},
		{Name: "list", Func: List},
		{Name: "loading", Func: Loading},
		{Name: "max", Func: Max},
		{Name: "maxlength", Func: MaxLength},
		{Name: "method", Func: Method},
		{Name: "min", Func: Min},
		{Name: "minlength", Func: MinLength},
		{Name: "name", Func: Name},
		{Name: "pattern", Func: Pattern},
		{Name: "placeholder", Func: Placeholder},
		{Name: "poster", Func: Poster},
		{Name: "preload", Func: Preload},
		{Name: "rel", Func: Rel},
		{Name: "role", Func: Role},
		{Name: "rows", Func: Rows},
		{Name: "rowspan", Func: RowSpan},
		{Name: "slot", Func: SlotAttr},
		{Name: "src", Func: Src},
		{Name: "srcset", Func: SrcSet},
		{Name: "step", Func: Step},
		{Name: "style", Func: Style},
		{Name: "style", Func: StyleAttr},
		{Name: "tabindex", Func: TabIndex},
		{Name: "target", Func: Target},
		{Name: "title", Func: Title},
		{Name: "title", Func: TitleAttr},
		{Name: "type", Func: Type},
		{Name: "value", Func: Value},
		{Name: "width", Func: Width},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			n := g.El("div", test.Func("hat"))
			b := new(bytes.Buffer)
			n.Render(b)
			assert.Equal(t, fmt.Sprintf(`<div %v="hat"></div>`, test.Name), b.String())
		})
	}
}

func TestAria(t *testing.T) {
	t.Run("returns an attribute which name is prefixed with aria-", func(t *testing.T) {
		n := Aria("selected", "true")
		b := new(bytes.Buffer)
		n.Render(b)
		assert.Equal(t, ` aria-selected="true"`, b.String())
	})
}

func TestData(t *testing.T) {
	t.Run("returns an attribute which name is prefixed with data-", func(t *testing.T) {
		n := Data("id", "partyhat")
		b := new(bytes.Buffer)
		n.Render(b)
		assert.Equal(t, ` data-id="partyhat"`, b.String())

		n = DataAttr("id", "partyhat")
		b.Reset()
		n.Render(b)
		assert.Equal(t, ` data-id="partyhat"`, b.String())
	})
}
