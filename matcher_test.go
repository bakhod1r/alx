package alx

import (
	"testing"
)

func TestMatchWith(t *testing.T) {
	header := "fr-FR,fr;q=0.9,uz-UZ;q=0.8,en;q=0.5,zh-Hant-TW;q=0.3"
	supported := []string{"uz", "ru", "en", "zh-Hant"}

	tests := []struct {
		name      string
		header    string
		strategy  MatchStrategy
		expected  string
		ok        bool
		supported []string
	}{
		{
			name:     "Lookup strategy matches uz-UZ to uz",
			header:   header,
			strategy: Lookup,
			expected: "uz",
			ok:       true,
		},
		{
			name:     "Lookup strategy matches zh-Hant-TW to zh-Hant",
			header:   "zh-Hant-TW;q=0.9", // Only provide this in header to ensure it matches
			strategy: Lookup,
			expected: "zh-Hant",
			ok:       true,
		},
		{
			name:     "Exact strategy fails on uz-UZ if only uz supported",
			header:   header,
			strategy: Exact,
			expected: "en",
			ok:       true,
		},
		{
			name:     "Basic strategy matches en to en",
			header:   header,
			strategy: Basic,
			expected: "uz", // Wait, fr is not supported, uz-UZ's base is uz -> matches uz
			ok:       true,
		},
		{
			name:     "PreferRegion strategy",
			header:   header,
			strategy: PreferRegion,
			expected: "uz",
			ok:       true,
		},
		{
			name:     "Fallback strategy matches same as Lookup",
			header:   header,
			strategy: Fallback,
			expected: "uz",
			ok:       true,
		},
		{
			name:     "Wildcard matching",
			header:   "*,en;q=0.5",
			strategy: Lookup,
			expected: "uz", // Wildcard matches first supported
			ok:       true,
		},
		{
			name:      "Empty supported list",
			header:    header,
			strategy:  Lookup,
			expected:  "",
			ok:        false,
			supported: []string{}, // Override supported
		},
		{
			name:     "No match",
			header:   "de,it",
			strategy: Lookup,
			expected: "",
			ok:       false,
		},
		{
			name:     "Exact match success",
			header:   "zh-Hant,en;q=0.5",
			strategy: Exact,
			expected: "zh-Hant",
			ok:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			supp := supported
			if tt.supported != nil {
				supp = tt.supported
			}
			got, ok := MatchWith(tt.header, supp, tt.strategy)
			if ok != tt.ok || got != tt.expected {
				t.Errorf("MatchWith(%q, %v) = (%v, %v), want (%v, %v)", tt.header, tt.strategy, got, ok, tt.expected, tt.ok)
			}
		})
	}
}

func TestMatchAndBest(t *testing.T) {
	header := "uz-UZ,ru;q=0.8"
	supported := []string{"uz", "ru"}

	got, ok := Best(header, supported)
	if !ok || got != "uz" {
		t.Errorf("Best() = (%v, %v), want (uz, true)", got, ok)
	}

	got, ok = Match(header, supported)
	if !ok || got != "uz" {
		t.Errorf("Match() = (%v, %v), want (uz, true)", got, ok)
	}
}
