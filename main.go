package main

import (
	"log"
	"os"
	"strconv"
	"system-monitor-backend/db"
	"system-monitor-backend/models"

	"system-monitor-backend/routers"

	"github.com/joho/godotenv"
)

func main() {
	// Load config file into environment variables
	err := godotenv.Load()
	panicIfErr(err)
	logFileName := os.Getenv("LOG_FILE")
	if logFileName == "" {
		logFileName = "server.log"
	}
	// Set up logger
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	panicIfErr(err)
	log.SetOutput(file)

	// Set up db
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	panicIfErr(err)

	dbConfig := db.Config{
		DbHost:     os.Getenv("DB_HOST"),
		Port:       dbPort,
		SslMode:    os.Getenv("DB_SSL_MODE"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
	database := db.ConnectToDb(dbConfig)
	var userStore *models.UserStore = &database
	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	panicIfErr(err)
	routers.Run(userStore, serverPort)
}

func panicIfErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
