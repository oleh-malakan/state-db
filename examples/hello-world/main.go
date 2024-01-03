package main

import (
	"log"

	db "github.com/oleh-malakan/core-db"
)

func main() {
	entity, err := db.Open("./hello-world.entity")
	if err != nil {
		log.Fatal(err)
	}
	defer entity.Close()

	for entity.Next() {
		var greeting string
		err = entity.Data("greeting").Scan(&greeting)
		if err != nil {
			log.Print(err)

			return
		}
	}
}
