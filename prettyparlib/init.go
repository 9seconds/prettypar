package prettyparlib

import "regexp"

var (
	regexpEmptyLine           = regexp.MustCompile(`^\s*$`)
	regexpOrderedListPrefix   = regexp.MustCompile(`^\s*((?:\d+|[a-z])\.)\s+`)
	regexpUnorderedListPrefix = regexp.MustCompile(`^\s*([-\*]+)\s+`)
)
