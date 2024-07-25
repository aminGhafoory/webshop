package models

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type PostgresConfig struct {
	DBurl string
}

func Open(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.Fatal("Can not connect to the DataBase")
		return nil, fmt.Errorf("open : %w", err)
	}

	return db, nil
}

func Migrate(db *sql.DB, dir string) error {

	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate:%w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate : %w", err)
	}

	return nil

}

func MigrateFS(db *sql.DB, fs fs.FS) error {
	goose.SetBaseFS(fs)

	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, ".")

}
