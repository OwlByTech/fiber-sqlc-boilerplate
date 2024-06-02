package database

import (
	"fmt"
	"log"
	"optitech/database/seeders"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
)

func Seeder(arg string) error {
	db, err := Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	repository.Queries = *sq.New(db)

	// NOTE: here we provide all the seeders with the same
	// structure like client
	seeders := []seeders.Seeder{
		&seeders.ClientSeeder{},
	}

	switch arg {
	case "up":
		if err := seederUp(seeders); err != nil {
			return err
		}

	case "down":
		if err := seederDown(seeders); err != nil {
			return err
		}

	default:
		return fmt.Errorf("You must provide up or down arguments")
	}

	return nil
}

func seederUp(seeders []seeders.Seeder) error {
	for _, seeder := range seeders {
		if err := seeder.Up(); err != nil {
			// TODO: add rollback
			return err
		}
	}

	return nil
}

func seederDown(seeders []seeders.Seeder) error {
	for _, seeder := range seeders {
		if err := seeder.Down(); err != nil {
			// TODO: add rollback
			return err
		}
	}

	return nil
}
