package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var m *migrate.Migrate

func Migrate(arg string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting current directory: %v", err)
	}

	p := fmt.Sprintf("file://%s", filepath.ToSlash(filepath.Join(wd, "database", "schemas")))

	db, err := Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalf("%v", err)
	}

	m, err = migrate.NewWithDatabaseInstance(
		p,
		"postgres", driver)

	if err != nil {
		log.Fatalf("%v", err)
	}

	switch arg {
	case "up":
		migrateUp()
	case "down":
		migrateDown()
	default:
		return fmt.Errorf("You must provide up or down argument")
	}

	return nil
}

func migrateUp() {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("%v", err)
	}
	log.Print("Migration up succefully")
}

func migrateDown() {
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("%v", err)
	}
	log.Print("Migration down succefully")
}
