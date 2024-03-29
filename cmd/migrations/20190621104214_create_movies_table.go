package main

import (
	"github.com/go-pg/pg/orm"
	"github.com/robinjoseph08/go-pg-migrations"
)

func init() {
	up := func(db orm.DB) error {
		_, err := db.Exec(`
			CREATE TABLE movies (
					id SERIAL PRIMARY KEY,
					title TEXT NOT NULL,
					release_date TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'utc')
			)
		`)
		return err
	}

	down := func(db orm.DB) error {
		_, err := db.Exec("DROP TABLE movies")
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20190621104214_create_movies_table", up, down, opts)
}
