// Package fiber adapts github.com/bakhod1r/alx to Fiber, exposing FromRequest and
// Languages for the framework's request context.
package fiber

import (
	"github.com/bakhod1r/alx"
	"github.com/gofiber/fiber/v2"
)

// FromRequest parses the language natively from the Fiber request headers.
func FromRequest(c *fiber.Ctx) alx.Language {
	langs := Languages(c)
	if len(langs) > 0 {
		return langs[0]
	}
	return alx.Language{}
}

// Languages extracts and parses all languages natively from the Fiber request headers.
func Languages(c *fiber.Ctx) []alx.Language {
	for _, key := range []string{"Accept-Language", "Language", "Lang"} {
		if val := string(c.Get(key)); val != "" {
			return alx.ParseSorted(val)
		}
	}
	return nil
}
