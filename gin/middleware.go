// Package gin adapts github.com/bakhod1r/alx to Gin, exposing FromRequest and
// Languages for the framework's request context.
package gin

import (
	"github.com/bakhod1r/alx"
	"github.com/gin-gonic/gin"
)

// FromRequest parses the language natively from the Gin request headers.
func FromRequest(c *gin.Context) alx.Language {
	return alx.FromRequest(c.Request)
}

// Languages extracts and parses all languages natively from the Gin request headers.
func Languages(c *gin.Context) []alx.Language {
	return alx.Languages(c.Request)
}
