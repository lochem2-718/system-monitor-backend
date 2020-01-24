package main

import (
	"log"
	"os"
	"strconv"
	"system-monitor-backend/db"

	"system-monitor-backend/config"
	"system-monitor-backend/routers"
)

func main() {
	// Load config
	config, err := config.NewConfig()
	panicIfErr(err)

	logFileName := config.GetSetting("LOG_FILE")
	if logFileName == "" {
		logFileName = "server.log"
	}
	// Set up logger
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	panicIfErr(err)
	log.SetOutput(file)

	// Set up db
	dbPort, err := strconv.Atoi(config.GetSetting("DB_PORT"))
	panicIfErr(err)

	dbConfig := db.Config{
		DbHost:     config.GetSetting("DB_HOST"),
		Port:       dbPort,
		SslMode:    config.GetSetting("DB_SSL_MODE"),
		DbName:     config.GetSetting("DB_NAME"),
		DbUser:     config.GetSetting("DB_USER"),
		DbPassword: config.GetSetting("DB_PASSWORD"),
	}
	database := db.ConnectToDb(dbConfig)
	serverPort, err := strconv.Atoi(config.GetSetting("SERVER_PORT"))
	panicIfErr(err)
	routers.Run(database, serverPort)
}

func panicIfErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
