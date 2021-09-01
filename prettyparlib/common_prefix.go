package prettyparlib

func getLinesCommonPrefix(lines []string, options FormatOptions) string {
	if len(lines) == 0 {
		return ""
	}

	prefix := lines[0]

	for i := 1; i < len(lines) && prefix != ""; i++ {
		prefix = getCommonPrefix(prefix, lines[i])
	}

	return options.ParagraphPrefixes.FindString(prefix)
}

func getCommonPrefix(one, another string) string {
	minStr := []rune(one)
	maxStr := []rune(another)

	if len(minStr) > len(maxStr) {
		minStr, maxStr = maxStr, minStr
	}

	for i, v := range minStr {
		if v != maxStr[i] {
			return string(minStr[:i])
		}
	}

	return string(minStr)
}
