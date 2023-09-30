package config

import (
	"database/sql"
	"fmt"
	"golang-crud/helper"

	_ "github.com/lib/pq" // Postgres goalng driver
	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgress"
	password = "hbjk"
	dbName   = "test"
)

func databaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbName=%s", host, port, user, password, dbName)
	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfEror(err)
	log.Info().Msg("database connected")
	return db
}
