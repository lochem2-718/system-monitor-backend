package models

// User -- A type representing a user
type User struct {
	Username  string // Primary key
	Password  string
	FirstName string
	LastName  string
}

// UserStore -- An interface to represent a database store for users
type UserStore interface {
	CreateUser(username string) User
	ReadUser(user User) Result
	UpdateUser(user User) Result
	DeleteUser(username string) Result
}
