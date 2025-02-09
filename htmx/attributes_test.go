package htmx_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	g "github.com/fuguohong1024/gomponents"
	. "github.com/fuguohong1024/gomponents/html"

	hx "github.com/fuguohong1024/gomponents/htmx"
	"github.com/fuguohong1024/gomponents/internal/assert"
)

func TestAttributes(t *testing.T) {
	cases := map[string]func(string) g.Node{
		"boost":        hx.Boost,
		"get":          hx.Get,
		"post":         hx.Post,
		"push-url":     hx.PushURL,
		"select":       hx.Select,
		"select-oob":   hx.SelectOOB,
		"swap":         hx.Swap,
		"swap-oob":     hx.SwapOOB,
		"target":       hx.Target,
		"trigger":      hx.Trigger,
		"vals":         hx.Vals,
		"confirm":      hx.Confirm,
		"delete":       hx.Delete,
		"disable":      hx.Disable,
		"disabled-elt": hx.DisabledElt,
		"disinherit":   hx.Disinherit,
		"encoding":     hx.Encoding,
		"ext":          hx.Ext,
		"headers":      hx.Headers,
		"history":      hx.History,
		"history-elt":  hx.HistoryElt,
		"include":      hx.Include,
		"indicator":    hx.Indicator,
		"params":       hx.Params,
		"patch":        hx.Patch,
		"preserve":     hx.Preserve,
		"prompt":       hx.Prompt,
		"put":          hx.Put,
		"replace-url":  hx.ReplaceURL,
		"request":      hx.Request,
		"sync":         hx.Sync,
		"validate":     hx.Validate,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output hx-%v="hat"`, name), func(t *testing.T) {
			n := g.El("div", fn("hat"))
			b := new(bytes.Buffer)
			n.Render(b)
			assert.Equal(t, fmt.Sprintf(`<div hx-%v="hat"></div>`, name), b.String())
		})
	}
}

func ExampleGet() {
	n := Button(hx.Post("/clicked"), hx.Swap("outerHTML"))
	_ = n.Render(os.Stdout)
	// Output: <button hx-post="/clicked" hx-swap="outerHTML"></button>
}

func TestOn(t *testing.T) {
	t.Run(`should output hx-on:click="alert('hat')"`, func(t *testing.T) {
		n := g.El("div", hx.On("click", "alert('hat')"))
		b := new(bytes.Buffer)
		n.Render(b)
		assert.Equal(t, `<div hx-on:click="alert('hat')"></div>`, b.String())
	})

	t.Run(`should output hx-on::before-request="alert('hat')"`, func(t *testing.T) {
		n := g.El("div", hx.On(":before-request", "alert('hat')"))
		b := new(bytes.Buffer)
		n.Render(b)
		assert.Equal(t, `<div hx-on::before-request="alert('hat')"></div>`, b.String())
	})
}
