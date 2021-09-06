package prettyparlib

import (
	"strings"
	"unicode/utf8"
)

type textBlock []string

func (t textBlock) String() string {
	return strings.Join(t, "\n")
}

func (t textBlock) Empty() bool {
	return len(t) == 0
}

func (t textBlock) Indent(firstPrefix, otherPrefix string) textBlock {
	t = t.makeCopy()

	if !t.Empty() {
		t[0] = firstPrefix + t[0]
	}

	for i := 1; i < len(t); i++ {
		t[i] = otherPrefix + t[i]
	}

	return t
}

func (t textBlock) Dedent(firstPrefix, otherPrefix string) textBlock {
	t = t.makeCopy()

	if !t.Empty() {
		t[0] = strings.TrimPrefix(t[0], firstPrefix)
	}

	for i := 1; i < len(t); i++ {
		t[i] = strings.TrimPrefix(t[i], otherPrefix)
	}

	return t
}

func (t textBlock) Wrap(lineLength int) textBlock {
	newBlock := textBlock{}
	currentLine := strings.Builder{}

	for _, word := range strings.Fields(t.String()) {
		wordLength := utf8.RuneCountInString(word)
		currentLineLength := utf8.RuneCountInString(currentLine.String())

		if currentLineLength > 0 && currentLineLength+wordLength >= lineLength {
			newBlock = append(newBlock, currentLine.String())
			currentLine.Reset()
		}

		if currentLine.Len() > 0 {
			currentLine.WriteRune(' ')
		}

		currentLine.WriteString(word)
	}

	if currentLine.Len() > 0 {
		newBlock = append(newBlock, currentLine.String())
	}

	if newBlock.Empty() {
		newBlock = t.makeCopy()
	}

	return newBlock
}

func (t textBlock) Map(mapper func(s string) string) textBlock {
	t = t.makeCopy()

	for i, v := range t {
		t[i] = mapper(v)
	}

	return t
}

func (t textBlock) Split() textBlocks {
	blocks := make(textBlocks, 0, len(t))
	next := textBlock{}

	for _, line := range t {
		if !regexpEmptyLine.MatchString(line) {
			next = append(next, line)

			continue
		}

		if !next.Empty() {
			blocks = append(blocks, next)
			next = textBlock{}
		}

		blocks = append(blocks, textBlock{line})
	}

	if !next.Empty() {
		blocks = append(blocks, next)
	}

	return blocks
}

func (t *textBlock) makeCopy() textBlock {
	newT := make(textBlock, len(*t))

	copy(newT, *t)

	return newT
}
