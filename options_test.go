package alx

import (
	"reflect"
	"testing"
)

func TestOptions(t *testing.T) {
	resolver := NewResolver(
		WithSupported("uz", "ru", "en"),
		WithDefault("uz"),
		WithFallback("uz-Cyrl", "uz"),
		WithURL(),
		WithQuery("lang"),
		WithCookie("locale"),
		WithHeader(),
	)

	expectedSupported := []string{"uz", "ru", "en"}
	if !reflect.DeepEqual(resolver.supported, expectedSupported) {
		t.Errorf("expected supported %v, got %v", expectedSupported, resolver.supported)
	}

	if resolver.defaultLang != "uz" {
		t.Errorf("expected default uz, got %v", resolver.defaultLang)
	}

	if resolver.customFallbacks["uz-Cyrl"] != "uz" {
		t.Errorf("expected custom fallback uz for uz-Cyrl, got %v", resolver.customFallbacks["uz-Cyrl"])
	}

	expectedExtractors := 4
	if len(resolver.extractors) != expectedExtractors {
		t.Errorf("expected %d extractors, got %d", expectedExtractors, len(resolver.extractors))
	}
}
