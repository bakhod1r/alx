// Package iris adapts github.com/bakhod1r/alx to Iris, exposing FromRequest and
// Languages for the framework's request context.
package iris

import (
	"github.com/bakhod1r/alx"
	"github.com/kataras/iris/v12"
)

// FromRequest parses the language natively from the Iris request headers.
func FromRequest(c iris.Context) alx.Language {
	return alx.FromRequest(c.Request())
}

// Languages extracts and parses all languages natively from the Iris request headers.
func Languages(c iris.Context) []alx.Language {
	return alx.Languages(c.Request())
}
