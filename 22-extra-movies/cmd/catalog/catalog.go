package main

import (
	"go-core/22-extra-movies/pkg/api"
	"go-core/22-extra-movies/pkg/storage/mongo"
	"log"
)

func main() {
	db, err := mongo.New()
	if err != nil {
		log.Fatal(err)
	}
	err = db.ImportMovies()
	if err != nil {
		log.Fatal(err)
	}
	api := api.New(db)
	api.Run(":8080")
}
