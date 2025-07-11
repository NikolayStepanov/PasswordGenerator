// Package handler contains HTTP request handlers for the application
package handler

import "context"

// Handler handles incoming HTTP requests
type Handler struct {
}

// GetPassword retrieves a password based
func (h *Handler) GetPassword(ctx context.Context) (string, error) {
	//TODO implement me
	panic("implement me")
}

// NewHandler creates and returns a new instance of Handler
func NewHandler() *Handler {
	return &Handler{}
}
