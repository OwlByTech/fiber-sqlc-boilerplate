package main

import (
	"log"
	"os"
	"strconv"

	"optitech/database"
	"optitech/internal/repository"
	"optitech/internal/router"
	sq "optitech/internal/sqlc"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	port, err := strconv.ParseUint(os.Getenv("PORT"), 10, 16)
	if err != nil {
		log.Fatalf("You must provide positive port number: %v", err)
	}

	repository.Queries = *sq.New(db)

	s := &router.Server{
		Port: uint16(port),
	}

	s.New()
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}
}
