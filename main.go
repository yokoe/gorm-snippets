package main

import "log"

func main() {
	s, err := findByID("model.Book")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println(s)
}
