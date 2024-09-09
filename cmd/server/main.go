package main

import (
	"log"
	"os"
	"strconv"

	"owlbytech/database"
	"owlbytech/internal/config"
	"owlbytech/internal/repository"
	"owlbytech/internal/router"
	sq "owlbytech/internal/sqlc"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)
	if err != nil {
		log.Fatalf("You must provide positive port number: %v", err)
	}

	repository.Queries = sq.New(db)

	s := &router.Server{
		Port: uint16(port),
	}

	s.New()
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}
}
