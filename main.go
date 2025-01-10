package main

import (
    "bytes"
    "flag"
    "fmt"
    "io"
    "os"
    "unicode"
)

type Counter int

// func check(e error) {
//     if e != nil {
//         panic(e)
//     }
// }

func main() {
    showLines := flag.Bool("l", false, "print the newline counts")
    showBytes := flag.Bool("c", false, "print the byte counts")
    countWords := flag.Bool("w", false, "print the number of words in the file")
    countCharacters := flag.Bool("m", false, "print the characters in the file")
    flag.Parse()

    var input io.Reader
    filename := ""
    args := flag.Args()

    if len(args) >= 1 {
        file, err := os.Open(args[0])
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
            os.Exit(1)
        }
        defer file.Close()
        input = file
        filename = args[0]
    } else {
        stat, _ := os.Stdin.Stat()
        if (stat.Mode() & os.ModeCharDevice) == 0 {
            input = os.Stdin
        } else {
            fmt.Println("No input provided")
            os.Exit(1)
        }
    }

    // Read all input into a buffer
    buf, err := io.ReadAll(input)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
        os.Exit(1)
    }
    
    noFlags := !*showLines && !*showBytes && !*countWords && !*countCharacters
    
    if noFlags {
        // Get line count
        lineCount := bytes.Count(buf, []byte{'\n'})
        if len(buf) > 0 && !bytes.HasSuffix(buf, []byte{'\n'}) {
            lineCount++
        }

        // Get word count
        wordCount := int64(0)
        inWord := false
        reader := bytes.NewReader(buf)
        
        for {
            r, _, err := reader.ReadRune()
            if err != nil {
                if err == io.EOF {
                    break
                }
                fmt.Fprintf(os.Stderr, "Error reading runes: %v\n", err)
                os.Exit(1)
            }
            
            if unicode.IsSpace(r) {
                inWord = false
            } else if !inWord {
                wordCount++
                inWord = true
            }
        }

        // Get byte count
        byteCount := len(buf)

        // Print in the format: lines words bytes filename
        if filename != "" {
            fmt.Printf("%8d %8d %8d %s\n", lineCount, wordCount, byteCount, filename)
        } else {
            fmt.Printf("%8d %8d %8d\n", lineCount, wordCount, byteCount)
        }
        return
    }

    if *showLines {
        lineCount := bytes.Count(buf, []byte{'\n'})
        if len(buf) > 0 && !bytes.HasSuffix(buf, []byte{'\n'}) {
            lineCount++
        }
        fmt.Printf("%8d\n", lineCount)
    }

    if *showBytes {
        byteCount := len(buf)
        fmt.Printf("%8d\n", byteCount)
    }

    if *countWords {
        wordCount := int64(0)
        inWord := false
        reader := bytes.NewReader(buf)
        
        for {
            r, _, err := reader.ReadRune()
            if err != nil {
                if err == io.EOF {
                    break
                }
                fmt.Fprintf(os.Stderr, "Error reading runes: %v\n", err)
                os.Exit(1)
            }
            
            if unicode.IsSpace(r) {
                inWord = false
            } else if !inWord {
                wordCount++
                inWord = true
            }
        }
        
        fmt.Printf("%8d  %s\n", wordCount, filename)
    }

    if *countCharacters {
        charCount := 0
        reader := bytes.NewReader(buf)
        
        for {
            _, _, err := reader.ReadRune()
            if err != nil {
                if err == io.EOF {
                    break
                }
                fmt.Fprintf(os.Stderr, "Error reading runes: %v\n", err)
                os.Exit(1)
            }
            charCount++
        }

        fmt.Printf("%8d\n", charCount)
    }
}

func (c *Counter) Write(p []byte) (n int, err error) {
    l := len(p)
    *c = Counter(int(*c) + l)
    return l, nil
}

