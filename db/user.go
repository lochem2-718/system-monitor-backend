package db

import (
	"errors"
	"system-monitor-backend/models"
)

// CreateUser -- A function to create a user in the database
func (db Database) CreateUser(username string) models.User {
	return models.User{}
}

// ReadUser -- A function to create a user in the database
func (db Database) ReadUser(user models.User) error {
	return errors.New("I'm not implemented yet")
}

// UpdateUser -- A function to create a user in the database
func (db Database) UpdateUser(user models.User) error {
	return errors.New("I'm not implemented yet")
}

// DeleteUser -- A function to create a user in the database
func (db Database) DeleteUser(username string) error {
	return errors.New("I'm not implemented yet")
}
