package seeders

import (
	"context"
	"log"
	"owlbytech/internal/repository"
	"time"

	sq "owlbytech/internal/sqlc"
)

type ClientSeeder struct{}

func (cs ClientSeeder) Up() error {
	ctx := context.Background()
	curTime := time.Now()
	client := sq.CreateClientParams{
		Email:     "developers@owlbytech.com",
		GivenName: "Developers",
		Password:  "password",
		Surname:   "Enjoy",
		CreatedAt: curTime,
	}

	_, err := repository.Queries.CreateClient(ctx, client)
	if err != nil {
		return err
	}

	log.Printf("Client Up seeder run successfully")
	return nil
}

func (cs ClientSeeder) Down() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllClients(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Client Down seeder run successfully")
	return nil
}
