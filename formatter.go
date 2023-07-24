package formatter

import (
	"fmt"
	"strings"
)

// FormatFloat - Formats a float to a string with 2 decimal places
//
//goland:noinspection ALL
func FormatFloat(f float64) string {
	return fmt.Sprintf("%.2f", f)
}

// Title - Formats a string to be title case
//
//goland:noinspection ALL
func Title(s string) string {
	return strings.Title(s)
}
