package alx

import "strings"

// Fallbacks generates a slice of language tags representing the fallback chain
// for a given tag. For example, "zh-Hant-TW" results in ["zh-Hant-TW", "zh-Hant", "zh"].
func Fallbacks(tag string) []string {
	if tag == "" || tag == "*" {
		return []string{tag}
	}

	parts := strings.Split(tag, "-")
	result := make([]string, 0, len(parts))

	for i := len(parts); i > 0; i-- {
		result = append(result, strings.Join(parts[:i], "-"))
	}

	return result
}
