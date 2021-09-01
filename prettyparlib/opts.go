package prettyparlib

import "regexp"

// FormatOptions contains options for Format function.
type FormatOptions struct {
	// LineLength defines a count of utf-8 runes in a line. Negative
	// value effectively means 0.
	LineLength int

	// A regular expression that detects a paragraph prefix.
	ParagraphPrefixes *regexp.Regexp
}
