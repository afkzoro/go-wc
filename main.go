package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode"

	"time"
)

type Counter int

func check(e error) {
    if e != nil {
        panic(e)
    }

}


func main() {
	showLines := flag.Bool("l", false, "print the newline counts")
    showBytes := flag.Bool("c", false, "print the byte counts")
    countWords := flag.Bool("w", false, "Count the number of words in the file")
    countCharacters := flag.Bool("m", false, "Count the characters in the file")
    flag.Parse()


	args := flag.Args()
    if len(args) < 1 {
        fmt.Println("Usage: program [-l|-c] filename")
        os.Exit(1)
    }

	filename := args[0]

	if *showBytes {
        file, err := os.Open(filename)
        check(err)
        defer file.Close()

        var c Counter
        _, err = io.Copy(&c, file)
        check(err)
        fmt.Printf("Bytes Count: %v   %v\n", c, filename)
    }

	if *showLines {
        buf, err := os.ReadFile(filename)
        check(err)

        lineCount := bytes.Count(buf, []byte{'\n'})
        if len(buf) > 0 && !bytes.HasSuffix(buf, []byte{'\n'}) {
            lineCount++
        }
        fmt.Printf("Line Count: %v\n", lineCount)
    }

    if *countWords {

        startTime := time.Now()
        file, err := os.Open(filename)
        check(err)

        reader := bufio.NewReaderSize(file, 16*1024)
        wordCount := int64(0)
        inWord := false
        
        for {
            r, _, err := reader.ReadRune()
            if err != nil {
                break
            }
            
            if unicode.IsSpace(r) {
                inWord = false
            } else if !inWord {
                wordCount++
                inWord = true
            }
        }
            elapsedTime := time.Since(startTime).Milliseconds()
            println("Time elapsed in milliseconds: ", elapsedTime)

            fmt.Printf("Word Count: %v  %v\n", wordCount, filename)
    }

    if *countCharacters {
        file, err := os.Open(filename)
        check(err)

        scanner := bufio.NewScanner(file)
        scanner.Split(bufio.ScanRunes)
        charCount := 0

        for scanner.Scan() {
            charCount++
        }

        fmt.Printf("Char Count: %v  %v\n", charCount, filename)

    }
}

// func isWordChar(b byte) bool {
//     return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
// }

func (c *Counter) Write (p []byte) (n int, err error) {
	l := len(p)
	*c = Counter(int(*c) + l)
	return l, nil
}
