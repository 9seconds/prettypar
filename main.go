package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/9seconds/prettypar/prettyparlib"
)

var version = "dev"

const (
	ExitCodeOk           = 0
	ExitCodeBadArg       = 1
	ExitCodeCannotFormat = 2
)

func main() {
	flag.Parse()

	if *VersionFlag {
		fmt.Println(version) // nolint: forbidigo
		os.Exit(ExitCodeOk)
	}

	words := strings.Fields(ParagraphPrefixesFlag)

	if len(words) == 0 {
		failf(ExitCodeCannotFormat, "paragraph prefixes are empty")
	}

	for i, v := range words {
		words[i] = regexp.QuoteMeta(v)
	}

	compiledRegexp, err := regexp.Compile(
		fmt.Sprintf(`^\s*(?:%s)?`, strings.Join(words, "|")),
	)
	if err != nil {
		failf(ExitCodeCannotFormat, "incorrect regexp %s: %s", ParagraphPrefixesFlag, err.Error())
	}

	opts := prettyparlib.FormatOptions{
		LineLength:        int(*LineLengthFlag),
		ParagraphPrefixes: compiledRegexp,
	}

	data, err := io.ReadAll(InputFileFlag)
	if err != nil {
		failf(ExitCodeCannotFormat, "cannot read stream: %s\n", err.Error())
	}

	fmt.Println(prettyparlib.Format(string(data), opts)) // nolint: forbidigo
}

func failf(exitCode int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(exitCode)
}
