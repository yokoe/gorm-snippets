package snippet_test

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/yokoe/gorm-snippets"
)

var update = flag.Bool("update", false, "update .golden files")

func TestFindFunctions(t *testing.T) {
	tests := []struct {
		name         string
		generateFunc func() (string, error)
	}{
		{"FindByID", func() (string, error) { return snippet.FindByID("model.Book") }},
		{"FindByParam", func() (string, error) { return snippet.FindByParam("model.Book", "Title", "string") }},
		{"BatchFindByID", func() (string, error) { return snippet.BatchFindByID("model.Book") }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			goldenFile := fmt.Sprintf("./testdata/%s.golden", tt.name)
			r, err := tt.generateFunc()
			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}
			b := []byte(r)
			if *update {
				if err := ioutil.WriteFile(goldenFile, b, 0644); err != nil {
					t.Fatalf("Failed to update golden file: %+v", err)
				}
			}
			expected, err := ioutil.ReadFile(goldenFile)
			if err != nil {
				t.Fatalf("Failed to read golden file: %+v", err)
			}
			if !bytes.Equal(expected, b) {
				t.Logf("Expect:\n%v", string(expected))
				t.Logf("Got:\n%v", string(b))
				t.Fatal("Result is not same to golden file.")
			}
		})
	}
}
