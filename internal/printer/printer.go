package printer

import (
	"fmt"

	"github.com/afkzoro/go-wc/internal/counter"
)

func PrintResults(counts *counter.Counts, filename string, flags *counter.Flags) {
	if flags.NoFlags() {
		if filename  != "" {
			fmt.Printf("%8d %8d %8d %s\n", counts.Lines, counts.Words, counts.Bytes, filename)
		} else {
			fmt.Printf("%8d %8d %8d\n", counts.Lines, counts.Words, counts.Bytes)
        }
        return
	}

	if flags.CountLines {
        fmt.Printf("%8d\n", counts.Lines)
    }
    if flags.CountBytes {
        fmt.Printf("%8d\n", counts.Bytes)
    }
    if flags.CountWords {
        fmt.Printf("%8d\n", counts.Words)
    }
    if flags.CountCharacters {
        fmt.Printf("%8d\n", counts.Characters)
    }
}