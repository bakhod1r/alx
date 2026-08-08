package alx

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	langs := []Language{
		{Tag: "en", Quality: 0.5},
		{Tag: "uz", Quality: 1.0},
		{Tag: "ru", Quality: 0.8},
		{Tag: "fr", Quality: 0.5}, // same quality as en, should remain after it in stable sort
	}

	expected := []Language{
		{Tag: "uz", Quality: 1.0},
		{Tag: "ru", Quality: 0.8},
		{Tag: "en", Quality: 0.5},
		{Tag: "fr", Quality: 0.5},
	}

	Sort(langs)

	if !reflect.DeepEqual(langs, expected) {
		t.Errorf("Sort() = %v, want %v", langs, expected)
	}
}
