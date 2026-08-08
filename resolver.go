package alx

import (
	"net/http"
)

// ResolveRequest resolves the best matching language for an HTTP request based
// on the configured extractors (URL, Query, Cookie, Header) and their priority.
func (r *Resolver) ResolveRequest(req *http.Request) Resolution {
	for _, ext := range r.extractors {
		val := ext.fn(req)
		if val == "" {
			continue
		}

		if res, ok := r.resolveSimple(val); ok {
			res.Source = ext.source
			return res
		}
	}

	// Fallback to default
	res, _ := r.resolveSimple(r.defaultLang)
	res.Source = SourceDefault
	return res
}

// Resolve resolves the best matching language for an Accept-Language header string.
func (r *Resolver) Resolve(header string) Resolution {
	if res, ok := r.resolveSimple(header); ok {
		res.Source = SourceHeader
		return res
	}
	res, _ := r.resolveSimple(r.defaultLang)
	res.Source = SourceDefault
	return res
}

func (r *Resolver) resolveSimple(val string) (Resolution, bool) {
	langs := ParseSorted(val)
	return r.matchLangs(langs)
}

func (r *Resolver) matchLangs(langs []Language) (Resolution, bool) {
	if len(r.supported) == 0 {
		return Resolution{}, false
	}

	normalizedSupported := make([]string, len(r.supported))
	for i, s := range r.supported {
		normalizedSupported[i] = Normalize(s)
	}

	for _, lang := range langs {
		if lang.Tag == "*" {
			return r.buildResolution(r.supported[0], lang, r.supported[0]), true
		}

		fallbacks := []string{lang.Tag}
		if custom, ok := r.customFallbacks[lang.Tag]; ok {
			fallbacks = append(fallbacks, custom)
		}

		standardFallbacks := Fallbacks(lang.Tag)
		if len(standardFallbacks) > 1 {
			fallbacks = append(fallbacks, standardFallbacks[1:]...)
		}

		for _, fallback := range fallbacks {
			for i, s := range normalizedSupported {
				if fallback == s {
					return r.buildResolution(r.supported[i], lang, fallback), true
				}
			}
		}
	}

	return Resolution{}, false
}

func (r *Resolver) buildResolution(supportedTag string, originalLang Language, matchedTag string) Resolution {
	// Parse the supported tag to populate properties like Language, Script, Region
	lang, _ := parseTag(supportedTag, 0)
	return Resolution{
		Tag:      supportedTag,
		Language: lang.Language,
		Script:   lang.Script,
		Region:   lang.Region,
		Quality:  originalLang.Quality,
		Matched:  matchedTag,
	}
}
