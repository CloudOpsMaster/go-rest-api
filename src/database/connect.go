package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func Connect(settings Settings) (err error) {

	sqlInfo := fmt.Sprintf("host=%s, port=%s, user=$s, password=%s, dbname=%s, sslmode=disable",
		settings.Host, settings.Port, settings.User, settings.Pass, settings.Name)

	DB, err = sql.Open("postgres", sqlInfo)
	if err != nil {
		return err
	}
	log.Printf("Database conection was created: %s  \n", sqlInfo)
	return nil
}
