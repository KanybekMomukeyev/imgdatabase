package models

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateDatabaseUp(path string, migrationsPath string) error {
	fmt.Printf("\n MigrateDatabaseUp migrationsPath=%s\n", migrationsPath)
	db, err := sql.Open("postgres", path)
	if err != nil {
		fmt.Printf("\n sql.Open ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
		return err
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Printf("\n postgres.WithInstance ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		fmt.Printf("\n migrate.NewWithDatabaseInstance ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
		return err
	}

	err = m.Up()

	if err != nil {
		fmt.Printf("\n err = m.Up() ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
		return err
	}

	defer m.Close()
	return err
}

func MigrateDatabaseDown(path string, migrationsPath string) error {
	db, err := sql.Open("postgres", path)
	if err != nil {
		fmt.Printf("\n sql.Open ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Printf("\n postgres.WithInstance ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		fmt.Printf("\n migrate.NewWithDatabaseInstance ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
		return err
	}

	err = m.Down()
	if err != nil {
		fmt.Printf("\n m.Down() ERROR\n")
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}

	defer m.Close()
	return err
}

// migrate create -ext sql -dir models/migrations create_user
