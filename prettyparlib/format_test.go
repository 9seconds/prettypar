package prettyparlib_test

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/9seconds/prettypar/prettyparlib"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	t.Parallel()

	inputDir := filepath.Join("testdata", "input")

	files, err := os.ReadDir(inputDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files { // nolint: paralleltest
		name := file.Name()

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			runFormatFileTest(t, name)
		})
	}

	randomInput := make([]byte, 1024)

	if _, err := rand.Read(randomInput); err != nil {
		t.Fatalf("cannot generate a random input: %v", err)
	}

	t.Run("random", func(t *testing.T) {
		t.Parallel()

		input := string(randomInput)

		assert.Equal(t, prettyparlib.Format(input, prettyparlib.FormatOptions{}), input)
	})
}

func runFormatFileTest(t *testing.T, name string) {
	t.Helper()

	inputFileName := filepath.Join("testdata", "input", name)
	outputFileName := filepath.Join("testdata", "output", name)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		t.Fatalf("cannot open an input file %s: %v", inputFileName, err)
	}

	defer inputFile.Close()

	inputFileReader := bufio.NewReader(inputFile)
	header := bytes.Buffer{}

	for {
		text, err := inputFileReader.ReadString('\n')
		if err != nil {
			t.Fatalf("cannot read input file %s: %v", inputFileName, err)
		}

		text = strings.TrimSpace(text)

		if strings.HasPrefix(text, "#") {
			continue
		}

		if text == "" {
			break
		}

		header.WriteString(text)
		header.WriteRune('\n')
	}

	opts := prettyparlib.FormatOptions{}
	if err := json.Unmarshal(header.Bytes(), &opts); err != nil {
		t.Fatalf("cannot read options: %v", err)
	}

	input, err := io.ReadAll(inputFileReader)
	if err != nil {
		t.Fatalf("cannot read input file: %v", err)
	}

	output, err := os.ReadFile(outputFileName)
	if err != nil {
		t.Fatalf("cannot read output file: %v", err)
	}

	assert.Equal(
		t,
		prettyparlib.Format(string(input), opts),
		string(output[:len(output)-1]),
	)
}
