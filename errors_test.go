package alx

import (
	"strings"
	"testing"
)

func TestParseError(t *testing.T) {
	err1 := &ParseError{Position: 5, Reason: "invalid token"}
	if !strings.Contains(err1.Error(), "invalid token at position 5") {
		t.Errorf("unexpected error string: %v", err1.Error())
	}

	err2 := &ParseError{Position: 10, Token: "abc", Reason: "bad token"}
	if !strings.Contains(err2.Error(), "bad token at position 10 (token: \"abc\")") {
		t.Errorf("unexpected error string: %v", err2.Error())
	}
}
