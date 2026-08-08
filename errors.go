package alx

import "fmt"

// ParseError describes an error that occurred while parsing an Accept-Language header.
type ParseError struct {
	Position int
	Token    string
	Reason   string
}

// Error implements the error interface.
func (e *ParseError) Error() string {
	if e.Token != "" {
		return fmt.Sprintf("%s at position %d (token: %q)", e.Reason, e.Position, e.Token)
	}
	return fmt.Sprintf("%s at position %d", e.Reason, e.Position)
}
