// Package fasthttp adapts github.com/bakhod1r/alx to fasthttp, exposing FromRequest and
// Languages for the framework's request context.
package fasthttp

import (
	"github.com/bakhod1r/alx"
	"github.com/valyala/fasthttp"
)

// FromRequest parses the language natively from the fasthttp request headers.
func FromRequest(ctx *fasthttp.RequestCtx) alx.Language {
	langs := Languages(ctx)
	if len(langs) > 0 {
		return langs[0]
	}
	return alx.Language{}
}

// Languages extracts and parses all languages natively from the fasthttp request headers.
func Languages(ctx *fasthttp.RequestCtx) []alx.Language {
	for _, key := range []string{"Accept-Language", "Language", "Lang"} {
		if val := string(ctx.Request.Header.Peek(key)); val != "" {
			return alx.ParseSorted(val)
		}
	}
	return nil
}
