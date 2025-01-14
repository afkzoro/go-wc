package printer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/afkzoro/go-wc/internal/counter"
)

func TestPrintResults(t *testing.T) {
    tests := []struct {
        name     string
        counts   *counter.Counts
        filename string
        flags    *counter.Flags
        expected string
    }{
        {
            name: "all counts with filename",
            counts: &counter.Counts{
                Lines: 3,
                Words: 20,
                Bytes: 100,
            },
            filename: "test.txt",
            flags:    &counter.Flags{},
            expected: "       3       20      100 test.txt\n",
        },
        {
            name: "all counts without filename",
            counts: &counter.Counts{
                Lines: 3,
                Words: 20,
                Bytes: 100,
            },
            filename: "",
            flags:    &counter.Flags{},
            expected: "       3       20      100\n",
        },
        {
            name: "only lines",
            counts: &counter.Counts{
                Lines: 3,
            },
            filename: "",
            flags:    &counter.Flags{CountLines: true},
            expected: "       3\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            old := os.Stdout
            r, w, _ := os.Pipe()
            os.Stdout = w

            // Create a buffer to store the output
            var buf bytes.Buffer

            PrintResults(tt.counts, tt.filename, tt.flags)

            // Close the write end of the pipe
            w.Close()

            // Read the output
            if _, err := io.Copy(&buf, r); err != nil {
                t.Error(err)
            }

            os.Stdout = old

            got := buf.String()
            if got != tt.expected {
                t.Errorf("PrintResults() = %q, want %q", got, tt.expected)
            }
        })
    }
}