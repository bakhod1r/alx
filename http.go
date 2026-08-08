package alx

import (
	"net/http"
	"strings"
)

// PrimaryLanguage returns the most preferred language from the Accept-Language header
// using a zero-allocation fast path. It simply returns the first valid language tag found
// without parsing full Q-values or sorting.
func PrimaryLanguage(header string) string {
	idx := strings.IndexAny(header, ",;")
	if idx != -1 {
		return strings.TrimSpace(header[:idx])
	}
	return strings.TrimSpace(header)
}

// SetContentLanguage sets the Content-Language HTTP response header.
func SetContentLanguage(w http.ResponseWriter, lang string) {
	w.Header().Set("Content-Language", lang)
}

// SetVary adds Accept-Language to the Vary HTTP response header.
func SetVary(w http.ResponseWriter) {
	w.Header().Add("Vary", "Accept-Language")
}

// Languages extracts and parses the language from the request, checking Accept-Language, Language, and Lang headers.
func Languages(r *http.Request) []Language {
	if r == nil || r.Header == nil {
		return nil
	}

	for _, key := range []string{"Accept-Language", "Language", "Lang"} {
		if val := r.Header.Get(key); val != "" {
			return ParseSorted(val)
		}
	}

	return nil
}

// FromRequest parses the Accept-Language header from the request and returns the highest quality Language.
func FromRequest(r *http.Request) Language {
	langs := Languages(r)
	if len(langs) > 0 {
		return langs[0]
	}
	return Language{}
}
