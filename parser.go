package alx

import (
	"strconv"
	"strings"
)

// ParseRaw parses the Accept-Language header and returns the languages in the exact
// order they appeared. It ignores malformed segments (lenient mode).
func ParseRaw(header string) []Language {
	if header == "" {
		return nil
	}

	var langs []Language
	parts := strings.Split(header, ",")
	pos := 0

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			pos += len(part) + 1
			continue
		}

		lang, err := parseTag(trimmed, pos+strings.Index(part, trimmed))
		if err == nil {
			langs = append(langs, lang)
		}
		pos += len(part) + 1
	}

	return langs
}

// Parse parses the Accept-Language header in lenient mode, preserving the original order.
func Parse(header string) []Language {
	return ParseRaw(header)
}

// ParseSorted parses the Accept-Language header, removes invalid segments,
// and returns the valid languages sorted by quality (descending).
func ParseSorted(header string) []Language {
	langs := ParseRaw(header)
	Sort(langs)
	return langs
}

// ParseStrict parses the Accept-Language header and returns an error if any segment is malformed.
func ParseStrict(header string) ([]Language, error) {
	if header == "" {
		return nil, nil
	}

	var langs []Language
	parts := strings.Split(header, ",")
	pos := 0

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			pos += len(part) + 1
			continue
		}

		lang, err := parseTag(trimmed, pos+strings.Index(part, trimmed))
		if err != nil {
			return nil, err
		}

		langs = append(langs, lang)
		pos += len(part) + 1
	}

	return langs, nil
}

// Validate checks if the Accept-Language header is completely valid.
// It returns an error if the header contains malformed elements.
func Validate(header string) error {
	_, err := ParseStrict(header)
	return err
}

func parseTag(token string, pos int) (Language, error) {
	// A token looks like "en-US" or "en-US;q=0.8" or "*;q=0"
	parts := strings.SplitN(token, ";", 2)

	tagStr := strings.TrimSpace(parts[0])
	if tagStr == "" {
		return Language{}, &ParseError{Position: pos, Token: token, Reason: "empty language tag"}
	}

	lang := Language{
		Quality: 1.0,
	}

	// Parse quality value if present
	if len(parts) > 1 {
		qPart := strings.TrimSpace(parts[1])
		if !strings.HasPrefix(qPart, "q=") && !strings.HasPrefix(qPart, "Q=") {
			return Language{}, &ParseError{Position: pos, Token: token, Reason: "invalid quality parameter format"}
		}
		qValStr := qPart[2:]
		qVal, err := strconv.ParseFloat(qValStr, 64)
		if err != nil || qVal < 0 || qVal > 1.0 {
			return Language{}, &ParseError{Position: pos, Token: token, Reason: "invalid q-value"}
		}
		lang.Quality = qVal
	}

	normalized := Canonical(tagStr)
	lang.Tag = normalized

	if normalized == "*" {
		lang.Language = "*"
		return lang, nil
	}

	subtags := strings.Split(normalized, "-")
	if len(subtags) > 0 {
		lang.Language = subtags[0]
	}

	// Extract Script (4 letters) and Region (2 letters or 3 digits/letters)
	for i := 1; i < len(subtags); i++ {
		sub := subtags[i]
		if len(sub) == 4 {
			lang.Script = sub
		} else if len(sub) == 2 || len(sub) == 3 {
			lang.Region = sub
		}
	}

	return lang, nil
}
