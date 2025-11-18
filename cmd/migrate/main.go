package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"

	configs "github.com/AboLojy/Carbon-Budget-Visualiser/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func buildMigrateURL() string {
	// Example format: postgres://user:password@host:port/dbname?sslmode=disable

	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(configs.Envs.DBUser, configs.Envs.DBPassword),
		Host:   fmt.Sprintf("%s:%s", configs.Envs.DBHost, configs.Envs.DBPort),
		Path:   configs.Envs.DBName,
	}
	// Add query parameters like sslmode
	q := u.Query()
	q.Set("sslmode", "disable")
	u.RawQuery = q.Encode()

	return u.String()
}

func main() {
	// The path to your migration files, using the 'file://' scheme
	flag.String("op", "up", "migration operation: up or down")
	flag.Parse()
	operation := flag.Lookup("op").Value.String()
	migrationsPath := "file://cmd/migrate/migrations"

	// The full database URL string
	databaseURL := buildMigrateURL()

	// 1. Create a new migration instance
	m, err := migrate.New(
		migrationsPath,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Error initializing migrate: %v", err)
	}

	// 2. Run all 'up' migrations
	log.Println("Applying database migrations...")
	switch operation {
	case "up":

		err = m.Up()

		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}

		if err == migrate.ErrNoChange {
			log.Println("Database already up-to-date. No migrations applied.")
		} else {
			log.Println("Database migrations applied successfully.")
		}
	case "down":
		log.Println("Reverting database migrations...")
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration down failed: %v", err)
		}
		if err == migrate.ErrNoChange {
			log.Println("No migrations to revert. Database is at the initial state.")
		} else {
			log.Println("Database migrations reverted successfully.")
		}
	default:
		log.Fatalf("Unknown operation: %s. Use 'up' or 'down'.", operation)
	}
	// For cleaning up and rolling back (optional, for testing):
	// err = m.Down()
	// ...

	// Check for any errors on the source or database drivers
	sourceErr, dbErr := m.Close()
	if sourceErr != nil {
		log.Printf("Warning: Error closing migration source: %v", sourceErr)
	}
	if dbErr != nil {
		log.Printf("Warning: Error closing migration database connection: %v", dbErr)
	}
}
