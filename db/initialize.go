package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	db "github.com/RageNeko26/online-store/db/sqlc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migration/*.sql
var fs embed.FS

func Migrate(source string) {
	d, err := iofs.New(fs, "migration")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering migrate error")
		}
	}()

	if err != nil {
		log.Fatal("Failed to find migration ", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, source)

	if err != nil {
		panic(err)
	}

	err = m.Up()

	if err != nil {
		panic(err)
	}

	fmt.Println("Migrate Schema is success...")
}

type DBCredentials struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewDBConnection(source *DBCredentials) *db.Queries {
	dbSource := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", source.User, source.Password, source.Host, source.Port, source.DBName)
	sqlDB, err := sql.Open("postgres", dbSource)

	defer Migrate(dbSource)

	if err != nil {
		log.Fatalf("Failed to connect database: %s", err.Error())
	}

	q := db.New(sqlDB)

	return q
}
