package snippet_test

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/iancoleman/strcase"
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
			goldenFile := fmt.Sprintf("./testdata/%s.golden", strcase.ToSnake(tt.name))

			r, err := tt.generateFunc()
			if err != nil {
				t.Fatalf("err = %v, want nil", err)
			}

			got := []byte(r)

			if *update {
				if err := ioutil.WriteFile(goldenFile, got, 0644); err != nil {
					t.Fatalf("Failed to update golden file: %+v", err)
				}
			}

			want, err := ioutil.ReadFile(goldenFile)
			if err != nil {
				t.Fatalf("Failed to read golden file: %+v", err)
			}
			if !bytes.Equal(got, want) {
				t.Logf("Got:\n%v", string(got))
				t.Logf("Want:\n%v", string(want))
				t.Fatal("Result is not same to golden file.")
			}
		})
	}
}
