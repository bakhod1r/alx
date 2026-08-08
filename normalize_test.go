package alx

import "testing"

func TestNormalize(t *testing.T) {
	tests := []struct {
		tag      string
		expected string
	}{
		{"*", "*"},
		{"en", "en"},
		{"EN-us", "en-US"},
		{"ZH-hant-tw", "zh-Hant-TW"},
		{"uz-uz", "uz-UZ"},
		{"sr-latn-rs", "sr-Latn-RS"},
		{"sl-rozaj-biske", "sl-rozaj-biske"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			if got := Normalize(tt.tag); got != tt.expected {
				t.Errorf("Normalize(%q) = %q, want %q", tt.tag, got, tt.expected)
			}
		})
	}
}

func TestCanonical(t *testing.T) {
	tests := []struct {
		tag      string
		expected string
	}{
		{"en_US", "en-US"},
		{"ZH_hant_tw", "zh-Hant-TW"},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			if got := Canonical(tt.tag); got != tt.expected {
				t.Errorf("Canonical(%q) = %q, want %q", tt.tag, got, tt.expected)
			}
		})
	}
}
