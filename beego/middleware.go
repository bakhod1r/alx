// Package beego adapts github.com/bakhod1r/alx to Beego, exposing FromRequest and
// Languages for the framework's request context.
package beego

import (
	"github.com/bakhod1r/alx"
	"github.com/beego/beego/v2/server/web/context"
)

// FromRequest parses the language natively from the Beego request headers.
func FromRequest(c *context.Context) alx.Language {
	return alx.FromRequest(c.Request)
}

// Languages extracts and parses all languages natively from the Beego request headers.
func Languages(c *context.Context) []alx.Language {
	return alx.Languages(c.Request)
}
