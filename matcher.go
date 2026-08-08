package alx

import "strings"

// MatchStrategy defines how parsed languages are matched against supported languages.
type MatchStrategy int

const (
	// Exact requires an exact match of the language tag.
	Exact MatchStrategy = iota
	// Basic matches based on the base language only (e.g., en-US matches en).
	Basic
	// Lookup uses standard fallback lookup (e.g., zh-Hant-TW -> zh-Hant -> zh).
	Lookup
	// Fallback is an alias for Lookup.
	Fallback
	// PreferRegion prefers matching the same region if possible, falls back to Lookup.
	PreferRegion
)

// MatchWith matches an Accept-Language header against a list of supported languages
// using the specified matching strategy.
func MatchWith(header string, supported []string, strategy MatchStrategy) (string, bool) {
	if len(supported) == 0 {
		return "", false
	}

	langs := ParseSorted(header)

	// Pre-process supported languages for faster lookup
	normalizedSupported := make([]string, len(supported))
	for i, s := range supported {
		normalizedSupported[i] = Normalize(s)
	}

	for _, lang := range langs {
		if lang.Tag == "*" {
			// A wildcard matches the first supported language
			return supported[0], true
		}

		for i, s := range normalizedSupported {
			switch strategy {
			case Exact:
				if lang.Tag == s {
					return supported[i], true
				}
			case Basic:
				supportedBase := strings.Split(s, "-")[0]
				if lang.Language == supportedBase {
					return supported[i], true
				}
			case Lookup, Fallback, PreferRegion:
				// PreferRegion is treated identically to Lookup in this basic implementation
				for _, fallback := range Fallbacks(lang.Tag) {
					if fallback == s {
						return supported[i], true
					}
				}
			}
		}
	}

	return "", false
}

// Match is an alias for Best, providing a simple way to get the best match.
func Match(header string, supported []string) (string, bool) {
	return Best(header, supported)
}

// Best matches an Accept-Language header against a list of supported languages
// using the default Lookup strategy, returning the best match.
func Best(header string, supported []string) (string, bool) {
	return MatchWith(header, supported, Lookup)
}
