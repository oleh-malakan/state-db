package main

import (
	"log"

	db "github.com/oleh-malakan/core-db"
)

func main() {
	entity, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer entity.Close()
}
