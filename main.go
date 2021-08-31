package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "dev"

func main() {
	flag.Parse()

	if *VersionFlag {
		fmt.Println(version) // nolint: forbidigo
		os.Exit(0)
	}

	fmt.Println(LineLengthFlag)
}
