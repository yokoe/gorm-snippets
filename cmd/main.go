package main

import (
	"log"

	snippet "github.com/yokoe/gorm-snippets"
)

func main() {
	s, err := snippet.FindByParam("model.Book", "UUID", "string")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println(s)
}
