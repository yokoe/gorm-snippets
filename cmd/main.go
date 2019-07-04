package main

import (
	"log"

	snippet "github.com/yokoe/gorm-snippets"
)

func main() {
	s, err := snippet.BatchFindByID("model.Book")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println(s)
}
