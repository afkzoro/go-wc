package reader

import (
	"fmt"
	"io"
	"os"
)

func GetInput(args []string) (io.Reader, string, error) {
	if len(args) >= 1 {
		file, err := os.Open(args[0])
		if err != nil {
			return nil, "", fmt.Errorf("error openinng file: %w", err)
		}
		return file, args[0], nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return os.Stdin, "", nil
	}

	return nil, "", fmt.Errorf("no input provided")
}