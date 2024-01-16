package config

import (
	"database/sql"
	"log"

	db "github.com/RageNeko26/online-store/db/sqlc"
)

func NewDBConnection(driver string, source string) *db.Queries {
	sqlDB, err := sql.Open(driver, source)

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	q := db.New(sqlDB)

	return q
}
