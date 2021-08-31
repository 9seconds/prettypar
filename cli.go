package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	DefaultLineLength = 79
	DefaultInputFile  = "-"
)

type LineLengthType uint

func (l *LineLengthType) String() string {
	return strconv.Itoa(int(*l))
}

func (l *LineLengthType) Set(value string) error {
	intVal, err := strconv.ParseUint(value, 10, 16) // nolint: gomnd
	if err != nil {
		return fmt.Errorf("line length has to be a positive number: %w", err)
	}

	if intVal == 0 {
		return errors.New("line number should be > 0")
	}

	*l = LineLengthType(intVal)

	return nil
}

type FileType struct {
	io.ReadCloser
}

func (f *FileType) String() string {
	return ""
}

func (f *FileType) Set(value string) error {
	*f = FileType{}

	switch value {
	case "", "-":
		f.ReadCloser = os.Stdout

		return nil
	}

	fp, err := os.Open(value)
	if err != nil {
		return fmt.Errorf("cannot open a file: %w", err)
	}

	f.ReadCloser = fp

	return nil
}

var (
	LineLengthFlag = func() *LineLengthType {
		defaultValue := os.Getenv("PRETTYPAR_LINELENGTH")
		if defaultValue == "" {
			defaultValue = strconv.Itoa(DefaultLineLength)
		}

		var flagVar LineLengthType

		if err := flagVar.Set(defaultValue); err != nil {
			fmt.Fprintf(os.Stderr, "incorrect PRETTYPAR_LINELENGTH value: %s", err.Error())
			os.Exit(1)
		}

		flag.Var(&flagVar, "line-length", "a length of line (envvar PRETTYPAR_LINELENGTH)")

		return &flagVar
	}()

	InputFileFlag = func() *FileType {
		var flagVar FileType

		if err := flagVar.Set(DefaultInputFile); err != nil {
			panic(err)
		}

		flag.Var(&flagVar, "input-file", "an input file to use. - means stdin (default -)")

		return &flagVar
	}()

	VersionFlag = func() *bool {
		var boolVar bool

		flag.BoolVar(&boolVar, "version", false, "print version")

		return &boolVar
	}()
)
