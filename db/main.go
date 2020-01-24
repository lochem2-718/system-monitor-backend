package db

import (
	"database/sql"
	"fmt"
	"log"

	// Drivers required for postgres
	_ "github.com/lib/pq"
)

// Database -- A struct representing a postgres database
type Database struct {
	postgresDb *sql.DB
}

// ConnectToDb -- A function to connection to the db
func ConnectToDb(dbConfig Config) Database {
	postgresDb, err := sql.Open("postgres", dbConfig.String())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return Database{
		postgresDb: postgresDb,
	}
}

// Config -- A struct representing the configuration of a database
type Config struct {
	DbHost     string
	Port       int
	SslMode    string
	DbName     string
	DbUser     string
	DbPassword string
}

func (cfg Config) String() string {
	return fmt.Sprintf("host=%s port=%d sslmode=%s dbname=%s user=%s "+
		"password=%s", cfg.DbHost, cfg.Port, cfg.SslMode, cfg.DbName,
		cfg.DbUser, cfg.DbPassword)
}
