package alx

import (
	"net/http"
	"net/url"
	"testing"
)

func TestResolver(t *testing.T) {
	resolver := NewResolver(
		WithSupported("uz-UZ", "ru-RU", "en-US"),
		WithDefault("en-US"),
		WithFallback("uz-Cyrl-UZ", "uz-UZ"),
		WithURL(),
		WithQuery("lang"),
		WithCookie("locale"),
		WithHeader(),
	)

	tests := []struct {
		name     string
		req      *http.Request
		expected Resolution
	}{
		{
			name: "URL path match",
			req: &http.Request{
				URL: &url.URL{Path: "/ru-RU/products"},
			},
			expected: Resolution{
				Tag:      "ru-RU",
				Language: "ru",
				Region:   "RU",
				Source:   SourceURL,
				Quality:  1.0,
				Matched:  "ru-RU",
			},
		},
		{
			name: "Query parameter match",
			req: &http.Request{
				URL: &url.URL{Path: "/products", RawQuery: "lang=uz-UZ"},
			},
			expected: Resolution{
				Tag:      "uz-UZ",
				Language: "uz",
				Region:   "UZ",
				Source:   SourceQuery,
				Quality:  1.0,
				Matched:  "uz-UZ",
			},
		},
		{
			name: "Cookie match",
			req: &http.Request{
				URL:    &url.URL{Path: "/products"},
				Header: http.Header{"Cookie": {"locale=en-US"}},
			},
			expected: Resolution{
				Tag:      "en-US",
				Language: "en",
				Region:   "US",
				Source:   SourceCookie,
				Quality:  1.0,
				Matched:  "en-US",
			},
		},
		{
			name: "Header match",
			req: &http.Request{
				URL:    &url.URL{Path: "/products"},
				Header: http.Header{"Accept-Language": {"ru-RU;q=0.9, uz-UZ;q=0.8"}},
			},
			expected: Resolution{
				Tag:      "ru-RU",
				Language: "ru",
				Region:   "RU",
				Source:   SourceHeader,
				Quality:  0.9,
				Matched:  "ru-RU",
			},
		},
		{
			name: "Custom fallback match from header",
			req: &http.Request{
				URL:    &url.URL{Path: "/products"},
				Header: http.Header{"Accept-Language": {"uz-Cyrl-UZ;q=1.0"}},
			},
			expected: Resolution{
				Tag:      "uz-UZ",
				Language: "uz",
				Region:   "UZ",
				Source:   SourceHeader,
				Quality:  1.0,
				Matched:  "uz-UZ",
			},
		},
		{
			name: "Fallback to default",
			req: &http.Request{
				URL: &url.URL{Path: "/products"},
			},
			expected: Resolution{
				Tag:      "en-US",
				Language: "en",
				Region:   "US",
				Source:   SourceDefault,
				Quality:  1.0,
				Matched:  "en-US",
			},
		},
		{
			name: "Resolve header function",
			req:  nil, // not used for Resolve(header)
			expected: Resolution{
				Tag:      "ru-RU",
				Language: "ru",
				Region:   "RU",
				Source:   SourceHeader,
				Quality:  0.9,
				Matched:  "ru-RU",
			},
		},
		{
			name: "Wildcard match",
			req: &http.Request{
				URL:    &url.URL{Path: "/products"},
				Header: http.Header{"Accept-Language": {"*;q=0.5"}},
			},
			expected: Resolution{
				Tag:      "uz-UZ",
				Language: "uz",
				Region:   "UZ",
				Source:   SourceHeader,
				Quality:  0.5,
				Matched:  "uz-UZ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Resolution
			if tt.name == "Resolve header function" {
				got = resolver.Resolve("ru-RU;q=0.9, uz-UZ;q=0.8")
			} else {
				got = resolver.ResolveRequest(tt.req)
			}

			if got != tt.expected {
				t.Errorf("Resolve() = %+v, want %+v", got, tt.expected)
			}
		})
	}
}

func TestResolveEmptySupported(t *testing.T) {
	resolver := NewResolver()
	res := resolver.ResolveRequest(&http.Request{})
	if res.Tag != "" {
		t.Errorf("Expected empty resolution, got %v", res)
	}
}

func TestOptionsCoverage(t *testing.T) {
	resolver := NewResolver(
		WithURL(),
		WithQuery("lang"),
		WithCookie("locale"),
		WithHeader(),
	)

	// test req == nil
	for _, ext := range resolver.extractors {
		if ext.fn(nil) != "" {
			t.Errorf("extractor %v returned non-empty for nil request", ext.source)
		}
	}

	req := &http.Request{} // no URL, no Header
	for _, ext := range resolver.extractors {
		if ext.fn(req) != "" {
			t.Errorf("extractor %v returned non-empty for empty request", ext.source)
		}
	}

	reqWithURL := &http.Request{URL: &url.URL{Path: "/"}}
	if resolver.extractors[0].fn(reqWithURL) != "" {
		t.Errorf("WithURL returned non-empty for root path")
	}

	reqWithBadCookie := &http.Request{Header: http.Header{"Cookie": {"invalid_cookie"}}}
	if resolver.extractors[2].fn(reqWithBadCookie) != "" {
		t.Errorf("WithCookie returned non-empty for invalid cookie")
	}

	reqWithEmptyHeader := &http.Request{Header: http.Header{}}
	if resolver.extractors[3].fn(reqWithEmptyHeader) != "" {
		t.Errorf("WithHeader returned non-empty for empty header map")
	}
}

func TestResolveFallback(t *testing.T) {
	resolver := NewResolver(WithSupported("en"), WithDefault("en"))
	res := resolver.Resolve("invalid-x-y-z") // Should fail matching, trigger default
	if res.Source != SourceDefault {
		t.Errorf("Expected SourceDefault, got %v", res.Source)
	}
}
