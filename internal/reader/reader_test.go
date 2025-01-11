package reader

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestGetInput(t *testing.T) {
	//Create a test file
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.txt")
	content := []byte("test content")

	if err := os.WriteFile(tempFile, content, 0666); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name	string
		args	[]string
		wantFile	bool
		wantErr	bool
		checkStdin	bool
	} {
		{
			name: "valid file",
			args: []string{tempFile},
			wantFile: true,
			wantErr: false,
		},
		{
			name: "non-existent file",
			args: []string{"non-existent.txt"},
			wantErr: true,
		},
		{
			name: "no args",
			args: []string{},
			wantErr: true,
		},
		{
			name: "no args with terminal stdin",
			args: []string{},
			checkStdin: false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader, filename, err :=  GetInput(tt.args)


            if (err != nil) != tt.wantErr {
                t.Errorf("GetInput() error = %v, wantErr %v", err, tt.wantErr)
                return
            }

            if tt.wantFile {
                if filename == "" {
                    t.Error("GetInput() expected filename, got empty string")
                }
                if reader == nil {
                    t.Error("GetInput() expected reader, got nil")
                }
                // Clean up
                if closer, ok := reader.(io.Closer); ok {
                    closer.Close()
                }
            }

            if tt.checkStdin {
                if reader != os.Stdin {
                    t.Error("GetInput() expected stdin reader")
                }
            }

		})
	}

}