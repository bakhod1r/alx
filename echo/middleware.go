// Package echo adapts github.com/bakhod1r/alx to Echo, exposing FromRequest and
// Languages for the framework's request context.
package echo

import (
	"github.com/bakhod1r/alx"
	"github.com/labstack/echo/v4"
)

// FromRequest parses the language natively from the Echo request headers.
func FromRequest(c echo.Context) alx.Language {
	return alx.FromRequest(c.Request())
}

// Languages extracts and parses all languages natively from the Echo request headers.
func Languages(c echo.Context) []alx.Language {
	return alx.Languages(c.Request())
}
