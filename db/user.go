package db

import (
	"system-monitor-backend/models"
)

// CreateUser -- A function to create a user in the database
func (db *Database) CreateUser(username string) models.User {
	return models.User{}
}

// ReadUser -- A function to create a user in the database
func (db *Database) ReadUser(user models.User) models.Result {
	return models.Failure
}

// UpdateUser -- A function to create a user in the database
func (db *Database) UpdateUser(user models.User) models.Result {
	return models.Failure
}

// DeleteUser -- A function to create a user in the database
func (db *Database) DeleteUser(username string) models.Result {
	return models.Failure
}
