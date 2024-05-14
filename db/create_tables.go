package db

import (
	"database/sql"
)

func CreateImageTable(database *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS image (
			id SERIAL PRIMARY KEY NOT NULL,
			name TEXT NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			deleted_at TIMESTAMP DEFAULT NULL,
			checksum bytea NOT NULL
	)
	`

	_, err := database.Exec(query)
	if err != nil {
			panic("failed to create image table")
	}
}