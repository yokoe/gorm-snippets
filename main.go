package main

import "log"

func main() {
	s, err := findByParam("model.Book", "UUID", "string")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println(s)
}
