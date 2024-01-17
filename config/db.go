package config

import (
	"database/sql"
	"fmt"
	"log"

	db "github.com/RageNeko26/online-store/db/sqlc"
	_ "github.com/lib/pq"
)

type DBCredentials struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewDBConnection(source *DBCredentials) *db.Queries {
	sqlDB, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", source.User, source.Password, source.Host, source.Port, source.DBName))

	if err != nil {
		log.Fatalf("Failed to connect database: %s", err.Error())
	}

	q := db.New(sqlDB)

	return q
}
