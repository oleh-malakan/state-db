package main

import (
	"log"

	db "github.com/oleh-malakan/state-db"
)

func main() {
	state := db.MakeState()
	defer state.Close()

	entity, err := state.Open("./hello-world.entity")
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
