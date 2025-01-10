package counter

import (
	"strings"
	"testing"

)

func TestCount(t *testing.T) {
	tests := []struct {
		name 		string
		input		string
		flags		*Flags
		expected	*Counts
		wantErr		bool 
	} {
	{
		name: "empty string",
		input: "",
		flags: &Flags{CountLines: true, CountBytes: true, CountWords: true},
		expected: &Counts{
			Lines: 0,
			Words: 0,
			Bytes: 0,
		},
	},
	{
		name: "single word",
		input: "hello",
		flags: &Flags{CountLines: true, CountBytes: true, CountWords: true},
		expected: &Counts{
			Lines: 1,
			Words: 1,
			Bytes: 5,
		},
	},
	{
		name: "multiple lines",
		input: "hello world\ntest line\nfinal",
		flags: &Flags{CountLines: true, CountBytes: true, CountWords: true},
		expected: &Counts{
			Lines: 3,
			Words: 5,
			Bytes: 27,
		},
	},
	{
		name: "multiple empty lines",
		input: "\n\n\n",
		flags: &Flags{CountLines: true, CountBytes: true, CountWords: true},
		expected: &Counts{
			Lines: 3,
			Words: 0,
			Bytes: 3,
		},
	},
	{
		name: "unicode characters",
		input: "hello 世界",
		flags: &Flags{CountLines: true, CountBytes: true, CountWords: true, CountCharacters: true},
		expected: &Counts{
			Lines: 1,
			Words: 2,
			Bytes: 12,
			Characters: 8,
		},
	},
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got, err := Count(reader, tt.flags)

			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got.Lines != tt.expected.Lines {
				t.Errorf("Count() lines = %v, want %v", got.Lines, tt.expected.Lines)
			}

			if got.Words != tt.expected.Words {
				t.Errorf("Count() words = %v, want %v", got.Words, tt.expected.Words)
			}

			if got.Bytes != tt.expected.Bytes {
				t.Errorf("Count() bytes = %v, want %v", got.Bytes, tt.expected.Bytes)
			}

			if tt.flags.CountCharacters && got.Characters != tt.expected.Characters {
                t.Errorf("Count() characters = %v, want %v", got.Characters, tt.expected.Characters)
            }


		})
	}
}

func TestFlags_NoFlags(t *testing.T) {
	tests := []struct {
		name		string
		flags 		*Flags
		expected	bool
 	} {
		{
			name: "no flags set",
			flags: &Flags{},
			expected: true,
		},
		{
			name: "line flags set",
			flags: &Flags{CountLines: true},
			expected: false,
		},
		{
			name: "all flags set",
			flags: &Flags{CountWords: true, CountBytes: true, CountLines: true, CountCharacters: true},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.flags.NoFlags(); got != tt.expected {
                t.Errorf("Flags.NoFlags() = %v, want %v", got, tt.expected)
            }
		})
	}
}