package models

import (
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func MigrateDatabaseUp(path string, migrationsPath string) error {
	db, err := sql.Open("postgres", path)
	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}

	err = m.Up()

	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}

	defer m.Close()
	return err
}

func MigrateDatabaseDown(path string, migrationsPath string) error {
	db, err := sql.Open("postgres", path)
	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}

	err = m.Down()
	if err != nil {
		log.WithFields(log.Fields{"omg": err}).Warn("migration error:")
	}

	defer m.Close()
	return err
}

// migrate create -ext sql -dir models/migrations create_user
