package alx

import (
	"fmt"
	"strconv"
)

// Language represents a parsed language tag and its associated quality value.
type Language struct {
	Tag      string  // The full language tag (e.g., "en-US")
	Language string  // The base language (e.g., "en")
	Script   string  // The script (e.g., "Hant")
	Region   string  // The region (e.g., "US")
	Quality  float64 // The quality value (q-factor), default is 1.0
}

// String returns the string representation of the language, including the q-value if it's not 1.0.
func (l Language) String() string {
	if l.Quality == 1.0 {
		return l.Tag
	}
	// For other values, format the float nicely.
	return fmt.Sprintf("%s;q=%s", l.Tag, strconv.FormatFloat(l.Quality, 'f', -1, 64))
}
