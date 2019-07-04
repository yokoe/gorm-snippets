package snippet_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"testing"

	"github.com/yokoe/gorm-snippets"
)

var update = flag.Bool("update", false, "update .golden files")

func TestFindByID(t *testing.T) {
	goldenFile := "./testdata/FindByID.golden"
	r, err := snippet.FindByID("model.Book")
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
}
