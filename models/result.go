package models

// Result -- An enum type to represent the result of an operation
type Result int8

const (
	// Success -- The enum value to represent a success result
	Success Result = 0
	// Failure -- The enum value to represent a failure result
	Failure Result = 1
)
