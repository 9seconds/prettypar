package prettyparlib

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// FormatOptions contains options for Format function.
type FormatOptions struct {
	// LineLength defines a count of utf-8 runes in a line. Negative
	// value effectively means 0.
	LineLength int

	// A regular expression that detects a paragraph prefix.
	ParagraphPrefixes *regexp.Regexp
}

func (f FormatOptions) MarshalJSON() ([]byte, error) {
	toJSON := struct {
		LineLength        int    `json:"line-length"`        // nolint: tagliatelle
		ParagraphPrefixes string `json:"paragraph-prefixes"` // nolint: tagliatelle
	}{
		LineLength:        f.LineLength,
		ParagraphPrefixes: f.ParagraphPrefixes.String(),
	}

	return json.Marshal(toJSON) // nolint: wrapcheck
}

func (f *FormatOptions) UnmarshalJSON(data []byte) error {
	fromJSON := struct {
		LineLength        int    `json:"line-length"`        // nolint: tagliatelle
		ParagraphPrefixes string `json:"paragraph-prefixes"` // nolint: tagliatelle
	}{}

	if err := json.Unmarshal(data, &fromJSON); err != nil {
		return err // nolint: wrapcheck
	}

	re, err := regexp.Compile(fromJSON.ParagraphPrefixes)
	if err != nil {
		return fmt.Errorf("cannot compile regular expression: %w", err)
	}

	f.LineLength = fromJSON.LineLength
	f.ParagraphPrefixes = re

	return nil
}
