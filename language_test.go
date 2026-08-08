package alx

import (
	"testing"
)

func TestLanguageString(t *testing.T) {
	tests := []struct {
		name     string
		lang     Language
		expected string
	}{
		{
			name:     "default quality",
			lang:     Language{Tag: "en-US", Quality: 1.0},
			expected: "en-US",
		},
		{
			name:     "custom quality",
			lang:     Language{Tag: "ru", Quality: 0.8},
			expected: "ru;q=0.8",
		},
		{
			name:     "zero quality",
			lang:     Language{Tag: "fr", Quality: 0},
			expected: "fr;q=0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lang.String(); got != tt.expected {
				t.Errorf("Language.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}
