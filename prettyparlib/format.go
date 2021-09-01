package prettyparlib

import (
	"bufio"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Format formats a given text based on provided options.
//
// Text is considered as a text block by lines. Lines are split by a '\n'
// (I do not have any intention to support '\r\n' and friends).
//
// Here is an anatomy of the text from Format's point of view:
//
//     # Some prefix
//     # lalala
//     #
//     #  1. blablabla
//     #     and continue
//     #  2. tututu
//
//  \--------------------/ <- paragraph
//        \--------------/ <- list block
//           \-----------/ <- text block
//
// Text may have many paragraphs. Each paragraph is delimited by 1+
// empty lines. A paragraph may have some common prefix, like
// comment etc.
//
// Comments are so common that this function manages them recursively.
//
// Each paragraph may be a list. List are made of this form:
//
//     1. N.
//     2. a.
//     3. -
//     4. *
//
// Nested lists are fine but they have to be a paragraph itself. Each list may
// have a multiple lines. Each line has to be adjusted to a first one. Like
// this:
//
//     1. Hahaha
//        Hahaha
//     2. lslsls
//
// Incorrect list is treated as a text block. Each item is treated as
// a text block and adjusted relative to its bullet.
//
// Text block is wrapped by a line length.
func Format(text string, options FormatOptions) string {
	if !utf8.ValidString(text) {
		return text
	}

	scanner := bufio.NewScanner(strings.NewReader(text))
	mainTextBlock := textBlock{}

	for scanner.Scan() {
		mainTextBlock = append(mainTextBlock, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return text
	}

	return mainTextBlock.
		Split().
		Map(func(b textBlock) textBlock {
			return formatParagraphBlock(b, options)
		}).
		Join().
		String()
}

func formatParagraphBlock(text textBlock, options FormatOptions) textBlock {
	prefix := getLinesCommonPrefix(text, options)
	options.LineLength -= len(prefix)

	formatFunc := formatListBlock

	if prefix != "" {
		formatFunc = formatParagraphBlock
	}

	return text.
		Dedent(prefix, prefix).
		Split().
		Map(func(b textBlock) textBlock {
			return formatFunc(b, options)
		}).
		Join().
		Indent(prefix, prefix)
}

func formatListBlock(text textBlock, options FormatOptions) textBlock {
	newBlock, ok := formatListBlockByRegexp(text, options, regexpOrderedListPrefix)

	if !ok {
		newBlock, ok = formatListBlockByRegexp(text, options, regexpUnorderedListPrefix)
	}

	if !ok {
		newBlock = formatTextBlock(text, options)
	}

	return newBlock
}

func formatListBlockByRegexp(text textBlock, options FormatOptions, re *regexp.Regexp) (textBlock, bool) {
	if text.Empty() {
		return text, false
	}

	firstMatch := re.FindStringSubmatch(text[0])
	if len(firstMatch) == 0 {
		return text, false
	}

	previousListPrefix := firstMatch[0]
	emptyPrefix := strings.Repeat(" ", len(previousListPrefix))
	options.LineLength -= len(previousListPrefix)

	textBlocks := textBlocks{}
	nextBlock := textBlock{}

	for _, line := range text {
		listPrefixMatch := re.FindStringSubmatch(line)

		if len(listPrefixMatch) == 0 {
			if !strings.HasPrefix(line, emptyPrefix) {
				return text, false
			}

			nextBlock = append(nextBlock, line)

			continue
		}

		if !nextBlock.Empty() {
			textBlocks = append(textBlocks,
				formatTextBlock(nextBlock, options).Indent(previousListPrefix, emptyPrefix))
			nextBlock = textBlock{}
		}

		previousListPrefix = listPrefixMatch[0]
		if len(previousListPrefix) != len(emptyPrefix) {
			return text, false
		}

		nextBlock = append(nextBlock, strings.TrimPrefix(line, previousListPrefix))
	}

	if !nextBlock.Empty() {
		textBlocks = append(textBlocks,
			formatTextBlock(nextBlock, options).Indent(previousListPrefix, emptyPrefix))
	}

	return textBlocks.Join(), true
}

func formatTextBlock(text textBlock, options FormatOptions) textBlock {
	return text.Map(strings.TrimSpace).Wrap(options.LineLength)
}
