package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil

	server := NewApiServer(":3000", store)
	server.Run()
}
