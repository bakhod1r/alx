package alx

import (
	"net/http"
	"strings"
)

// Source identifies where the language was resolved from.
type Source string

// The sources a language can be resolved from, in the order a Resolver is
// typically configured to consult them.
const (
	SourceURL     Source = "url"
	SourceQuery   Source = "query"
	SourceCookie  Source = "cookie"
	SourceHeader  Source = "header"
	SourceDefault Source = "default"
)

// Resolution represents the result of resolving a locale.
type Resolution struct {
	Tag      string
	Language string
	Script   string
	Region   string
	Source   Source
	Quality  float64
	Matched  string
}

// Resolver resolves the best language for an HTTP request or header.
type Resolver struct {
	supported       []string
	defaultLang     string
	customFallbacks map[string]string
	extractors      []extractor
}

type extractor struct {
	source Source
	fn     func(*http.Request) string
}

// ResolverOption configures the Resolver.
type ResolverOption func(*Resolver)

// NewResolver creates a new Resolver with the given options.
func NewResolver(opts ...ResolverOption) *Resolver {
	r := &Resolver{
		customFallbacks: make(map[string]string),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// WithSupported sets the list of supported languages.
func WithSupported(langs ...string) ResolverOption {
	return func(r *Resolver) {
		r.supported = langs
	}
}

// WithDefault sets the default language to use if no match is found.
func WithDefault(lang string) ResolverOption {
	return func(r *Resolver) {
		r.defaultLang = lang
	}
}

// WithFallback adds a custom fallback mapping.
func WithFallback(tag, fallback string) ResolverOption {
	return func(r *Resolver) {
		r.customFallbacks[tag] = fallback
	}
}

// WithURL enables extracting the language from the first path segment of the URL (e.g. /en/products).
func WithURL() ResolverOption {
	return func(r *Resolver) {
		r.extractors = append(r.extractors, extractor{
			source: SourceURL,
			fn: func(req *http.Request) string {
				if req == nil || req.URL == nil || req.URL.Path == "" {
					return ""
				}
				path := strings.TrimPrefix(req.URL.Path, "/")
				segments := strings.SplitN(path, "/", 2)
				return segments[0]
			},
		})
	}
}

// WithQuery enables extracting the language from a query parameter.
func WithQuery(param string) ResolverOption {
	return func(r *Resolver) {
		r.extractors = append(r.extractors, extractor{
			source: SourceQuery,
			fn: func(req *http.Request) string {
				if req == nil || req.URL == nil {
					return ""
				}
				return req.URL.Query().Get(param)
			},
		})
	}
}

// WithCookie enables extracting the language from a cookie.
func WithCookie(name string) ResolverOption {
	return func(r *Resolver) {
		r.extractors = append(r.extractors, extractor{
			source: SourceCookie,
			fn: func(req *http.Request) string {
				if req == nil {
					return ""
				}
				cookie, err := req.Cookie(name)
				if err == nil && cookie != nil {
					return cookie.Value
				}
				return ""
			},
		})
	}
}

// WithHeader enables extracting the language from HTTP headers.
// If no headers are specified, it defaults to "Accept-Language".
func WithHeader(headers ...string) ResolverOption {
	if len(headers) == 0 {
		headers = []string{"Accept-Language"}
	}
	return func(r *Resolver) {
		r.extractors = append(r.extractors, extractor{
			source: SourceHeader,
			fn: func(req *http.Request) string {
				if req == nil || req.Header == nil {
					return ""
				}
				for _, h := range headers {
					if val := req.Header.Get(h); val != "" {
						return val
					}
				}
				return ""
			},
		})
	}
}
