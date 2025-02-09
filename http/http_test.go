package http_test

import (
	"net/http"
	"testing"

	"github.com/fuguohong1024/gomponents/internal/assert"

	hxhttp "github.com/fuguohong1024/gomponents/http"
)

func TestBoolGetters(t *testing.T) {
	tests := map[string]func(http.Header) bool{
		"Boosted":                 hxhttp.IsBoosted,
		"History-Restore-Request": hxhttp.IsHistoryRestoreRequest,
		"Request":                 hxhttp.IsRequest,
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}

			v := fn(headers)
			assert.Equal(t, false, v)

			headers.Set("HX-"+name, "true")
			v = fn(headers)
			assert.Equal(t, true, v)
		})
	}
}

func TestValueGetters(t *testing.T) {
	tests := map[string]func(http.Header) string{
		"Current-URL":  hxhttp.GetCurrentURL,
		"Prompt":       hxhttp.GetPrompt,
		"Target":       hxhttp.GetTarget,
		"Trigger-Name": hxhttp.GetTriggerName,
		"Trigger":      hxhttp.GetTrigger,
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}

			v := fn(headers)
			assert.Equal(t, "", v)

			headers.Set("HX-"+name, "foo")
			v = fn(headers)
			assert.Equal(t, "foo", v)
		})
	}
}

func TestSetRefresh(t *testing.T) {
	headers := http.Header{}
	hxhttp.SetRefresh(headers)
	assert.Equal(t, "true", headers.Get("HX-Refresh"))
}

func TestSetters(t *testing.T) {
	tests := map[string]func(http.Header, string){
		"Location":             hxhttp.SetLocation,
		"Push-Url":             hxhttp.SetPushURL,
		"Redirect":             hxhttp.SetRedirect,
		"Replace-Url":          hxhttp.SetReplaceURL,
		"Reswap":               hxhttp.SetReswap,
		"Retarget":             hxhttp.SetRetarget,
		"Trigger":              hxhttp.SetTrigger,
		"Trigger-After-Settle": hxhttp.SetTriggerAfterSettle,
		"Trigger-After-Swap":   hxhttp.SetTriggerAfterSwap,
	}

	for name, fn := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			fn(headers, "foo")
			assert.Equal(t, "foo", headers.Get("HX-"+name))
		})
	}
}

func ExampleIsBoosted() {
	_ = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hxhttp.IsBoosted(r.Header) {
			// Boosted!
		}
	})
}

func ExampleSetRefresh() {
	_ = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hxhttp.SetRefresh(w.Header())
	})
}
