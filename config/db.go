package config

import (
	"database/sql"
	"log"

	db "github.com/RageNeko26/online-store/db/sqlc"
	_ "github.com/lib/pq"
)

func NewDBConnection(source string) *db.Queries {
	sqlDB, err := sql.Open("postgres", source)

	if err != nil {
		log.Fatalf("Failed to connect database: %s", err.Error())
	}

	q := db.New(sqlDB)

	return q
}
