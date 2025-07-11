// Package repository provides access to the application's data storage layer
package repository

// PasswordStorager defines the interface for password storage operation
type PasswordStorager interface {
}

// Repository aggregates all data storage interfaces used by the application
type Repository struct {
	Passwords PasswordStorager
}

// NewRepository initializes a new Repository with all required storage dependencies
func NewRepository(passwordStorager PasswordStorager) *Repository {
	return &Repository{Passwords: passwordStorager}
}
