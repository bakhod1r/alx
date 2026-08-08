package alx

import (
	"strings"
)

// Normalize normalizes a language tag to its canonical casing.
// It formats the language code to lowercase, script code to titlecase,
// and region code to uppercase. For example: "EN-us" -> "en-US", "ZH-hant-tw" -> "zh-Hant-TW".
func Normalize(tag string) string {
	if tag == "*" {
		return tag
	}

	parts := strings.Split(tag, "-")

	// First part is the primary language subtag (lowercase)
	parts[0] = strings.ToLower(parts[0])

	for i := 1; i < len(parts); i++ {
		part := parts[i]
		if len(part) == 4 {
			// Script subtag (Titlecase)
			parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
		} else if len(part) == 2 || len(part) == 3 {
			// Region subtag (Uppercase)
			parts[i] = strings.ToUpper(part)
		} else {
			// Other subtags (lowercase)
			parts[i] = strings.ToLower(part)
		}
	}

	return strings.Join(parts, "-")
}

// Canonical transforms the input tag into a canonical format by replacing
// underscores with dashes and then applying Normalization.
func Canonical(tag string) string {
	tag = strings.ReplaceAll(tag, "_", "-")
	return Normalize(tag)
}
