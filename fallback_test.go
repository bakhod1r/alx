package alx

import (
	"reflect"
	"testing"
)

func TestFallbacks(t *testing.T) {
	tests := []struct {
		tag      string
		expected []string
	}{
		{"zh-Hant-TW", []string{"zh-Hant-TW", "zh-Hant", "zh"}},
		{"en-US", []string{"en-US", "en"}},
		{"uz", []string{"uz"}},
		{"*", []string{"*"}},
		{"", []string{""}},
		{"sr-Latn-RS", []string{"sr-Latn-RS", "sr-Latn", "sr"}},
	}

	for _, tt := range tests {
		t.Run(tt.tag, func(t *testing.T) {
			if got := Fallbacks(tt.tag); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Fallbacks(%q) = %v, want %v", tt.tag, got, tt.expected)
			}
		})
	}
}
