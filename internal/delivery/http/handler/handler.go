// Package handler contains HTTP request handlers for the application
package handler

import (
	"context"

	"github.com/NikolayStepanov/PasswordGenerator/internal/domain/dto"
	"github.com/NikolayStepanov/PasswordGenerator/internal/service"
)

// Handler handles incoming HTTP requests
type Handler struct {
	passwordService service.Password
}

// GetNewPassword retrieves a password based
func (h *Handler) GetNewPassword(ctx context.Context, options dto.GeneratePasswordOptions) (string, error) {
	pass, err := h.passwordService.GetNewPassword(ctx, options)
	return pass, err
}

// NewHandler creates and returns a new instance of Handler
func NewHandler(password service.Password) *Handler {
	return &Handler{
		password,
	}
}
