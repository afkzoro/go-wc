package counter

import (
	"bytes"
	"flag"
	"io"
	"unicode"
)

type Flags struct {
	CountLines		bool
	CountBytes		bool
	CountWords		bool
	CountCharacters	bool
}

type Counts struct {
	Lines		int
	Bytes		int
	Words		int64
	Characters	int
}

func NewFlags() *Flags {
	f := &Flags{}

	flag.BoolVar(&f.CountLines, "l", false, "Counts the lines in a file")
	flag.BoolVar(&f.CountBytes, "c", false, "Counts the bytes in a file")
	flag.BoolVar(&f.CountWords, "w", false, "Count the words in a file")
	flag.BoolVar(&f.CountCharacters, "m", false, "Count the characters in a file")
	return f
}

func (f *Flags) Parse() {
	flag.Parse()
}

func (f *Flags) NoFlags() bool {
	return !f.CountLines && !f.CountBytes && !f.CountWords && !f.CountCharacters
}

func Count(input io.Reader, flags *Flags) (*Counts, error) {
	buf, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	counts := &Counts{}

	if flags.NoFlags() || flags.CountLines {
		counts.Lines = countLines(buf)
	}

	if flags.NoFlags() || flags.CountBytes {
		counts.Bytes = len(buf)
	}

	if flags.NoFlags() || flags.CountWords {
		counts.Words = countWords(buf)
	}

	if flags.CountCharacters {
		counts.Characters = countCharacters(buf)
	}

	return counts, nil
}

func countLines (buf []byte) int {
	lines := bytes.Count(buf, []byte{'\n'})
	if len(buf) > 0 && !bytes.HasSuffix(buf, []byte{'\n'}) {
		lines++
	}
	return lines
}


func countWords(buf []byte) int64 {
	var wordCount int64
	inWord := false
	reader := bytes.NewReader(buf)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			return wordCount
		}

		if unicode.IsSpace(r) {
			inWord = false
		} else if !inWord {
			wordCount++
			inWord = true
		}
	}

	return wordCount
}

func countCharacters(buf []byte) int {
	return len(bytes.Runes(buf))
}
