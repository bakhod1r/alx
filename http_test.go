package alx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPrimaryLanguage(t *testing.T) {
	tests := []struct {
		header   string
		expected string
	}{
		{"uz-UZ,ru;q=0.8,en;q=0.5", "uz-UZ"},
		{"ru;q=0.9", "ru"},
		{"en-US", "en-US"},
		{"  fr-FR  , en", "fr-FR"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.header, func(t *testing.T) {
			if got := PrimaryLanguage(tt.header); got != tt.expected {
				t.Errorf("PrimaryLanguage(%q) = %q, want %q", tt.header, got, tt.expected)
			}
		})
	}
}

func BenchmarkPrimaryLanguage(b *testing.B) {
	header := "uz-UZ,ru;q=0.8,en;q=0.5"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrimaryLanguage(header)
	}
}

func TestSetContentLanguage(t *testing.T) {
	w := httptest.NewRecorder()
	SetContentLanguage(w, "uz-UZ")

	if got := w.Header().Get("Content-Language"); got != "uz-UZ" {
		t.Errorf("SetContentLanguage() set %q, want %q", got, "uz-UZ")
	}
}

func TestSetVary(t *testing.T) {
	w := httptest.NewRecorder()
	SetVary(w)

	// httptest.NewRecorder adds to header map, which we can check
	vary := w.Header().Values("Vary")
	found := false
	for _, v := range vary {
		if v == "Accept-Language" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("SetVary() did not add Accept-Language to Vary header. Vary headers: %v", vary)
	}
}

func TestLanguages(t *testing.T) {
	req := &http.Request{Header: http.Header{"Accept-Language": {"uz-UZ, ru;q=0.8"}}}
	langs := Languages(req)

	if len(langs) != 2 || langs[0].Tag != "uz-UZ" || langs[1].Tag != "ru" {
		t.Errorf("Languages() returned unexpected result: %v", langs)
	}

	if len(Languages(nil)) != 0 {
		t.Errorf("Languages(nil) should return nil")
	}

	reqEmpty := &http.Request{Header: http.Header{}}
	if len(Languages(reqEmpty)) != 0 {
		t.Errorf("Languages(reqEmpty) should return nil")
	}
}

func TestFromRequest(t *testing.T) {
	req := &http.Request{Header: http.Header{"Accept-Language": {"uz-UZ, ru;q=0.8"}}}
	lang := FromRequest(req)

	if lang.Tag != "uz-UZ" {
		t.Errorf("FromRequest() = %v, want uz-UZ", lang.Tag)
	}

	emptyLang := FromRequest(nil)
	if emptyLang.Tag != "" {
		t.Errorf("FromRequest(nil) should return empty Language")
	}
}
