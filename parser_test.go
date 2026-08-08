package alx

import (
	"reflect"
	"testing"
)

func TestParseStrict(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    []Language
		wantErr bool
	}{
		{
			name:   "empty",
			header: "",
			want:   nil,
		},
		{
			name:   "single language",
			header: "en-US",
			want: []Language{
				{Tag: "en-US", Language: "en", Region: "US", Quality: 1.0},
			},
		},
		{
			name:   "with quality",
			header: "en-US;q=0.8",
			want: []Language{
				{Tag: "en-US", Language: "en", Region: "US", Quality: 0.8},
			},
		},
		{
			name:   "multiple languages",
			header: "fr-FR, fr;q=0.9, en;q=0.8",
			want: []Language{
				{Tag: "fr-FR", Language: "fr", Region: "FR", Quality: 1.0},
				{Tag: "fr", Language: "fr", Quality: 0.9},
				{Tag: "en", Language: "en", Quality: 0.8},
			},
		},
		{
			name:   "wildcard",
			header: "*;q=0.5",
			want: []Language{
				{Tag: "*", Language: "*", Quality: 0.5},
			},
		},
		{
			name:   "complex tag",
			header: "zh-Hant-TW",
			want: []Language{
				{Tag: "zh-Hant-TW", Language: "zh", Script: "Hant", Region: "TW", Quality: 1.0},
			},
		},
		{
			name:   "case normalization",
			header: "EN-us;q=1",
			want: []Language{
				{Tag: "en-US", Language: "en", Region: "US", Quality: 1.0},
			},
		},
		{
			name:   "extra spaces and commas",
			header: " en , ,, fr; q=0.5 ",
			want: []Language{
				{Tag: "en", Language: "en", Quality: 1.0},
				{Tag: "fr", Language: "fr", Quality: 0.5},
			},
		},
		{
			name:    "invalid quality",
			header:  "en;q=abc",
			wantErr: true,
		},
		{
			name:    "invalid quality out of bounds high",
			header:  "en;q=1.5",
			wantErr: true,
		},
		{
			name:    "invalid quality out of bounds low",
			header:  "en;q=-0.1",
			wantErr: true,
		},
		{
			name:    "invalid format",
			header:  "en;x=0.5",
			wantErr: true,
		},
		{
			name:    "empty tag with semi",
			header:  ";q=0.5",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStrict(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseStrict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	// Parse should ignore invalid tags
	header := "en-US, invalid;x, fr;q=0.5"
	want := []Language{
		{Tag: "en-US", Language: "en", Region: "US", Quality: 1.0},
		{Tag: "fr", Language: "fr", Quality: 0.5},
	}

	got := Parse(header)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Parse() = %v, want %v", got, want)
	}
}

func TestParseSorted(t *testing.T) {
	header := "en;q=0.5, uz;q=1, ru;q=0.8"
	want := []Language{
		{Tag: "uz", Language: "uz", Quality: 1.0},
		{Tag: "ru", Language: "ru", Quality: 0.8},
		{Tag: "en", Language: "en", Quality: 0.5},
	}

	got := ParseSorted(header)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseSorted() = %v, want %v", got, want)
	}
}

func TestValidate(t *testing.T) {
	if err := Validate("en-US, en;q=0.5"); err != nil {
		t.Errorf("Validate() unexpected error: %v", err)
	}
	if err := Validate("en;q=abc"); err == nil {
		t.Errorf("Validate() expected error")
	}
}

func TestParseRaw(t *testing.T) {
	if got := ParseRaw(""); got != nil {
		t.Errorf("ParseRaw empty expected nil, got %v", got)
	}

	// Test error skipping and empty segments
	got := ParseRaw("en-US, , invalid;x, fr")
	if len(got) != 2 || got[0].Tag != "en-US" || got[1].Tag != "fr" {
		t.Errorf("ParseRaw with error expected 2 valid langs, got %v", got)
	}
}
