package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/afkzoro/go-wc/internal/counter"
	"github.com/afkzoro/go-wc/internal/printer"
	"github.com/afkzoro/go-wc/internal/reader"
)

func main() {
	flags := counter.NewFlags()
	flags.Parse()

	input, filename, err := reader.GetInput(flag.Args())
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err2 != nil {
			return
		}
		os.Exit(1)
	}

	counts, err := counter.Count(input, flags)
	if err != nil {
		_, err2 := fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err2 != nil {
			return
		}
		os.Exit(1)
	}

	printer.PrintResults(counts, filename, flags)
}
