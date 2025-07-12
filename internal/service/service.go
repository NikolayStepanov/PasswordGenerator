// Package service contains business logic and service layer implementations
// that handle core application functionality
package service

import (
	"context"

	"github.com/NikolayStepanov/PasswordGenerator/internal/domain/dto"
)

type Password interface {
	GetNewPassword(ctx context.Context, options dto.GeneratePasswordOptions) (string, error)
}

// Services holds application dependencies
type Services struct {
	Password Password
}

// NewServices creates and returns a new Services instance
func NewServices(passwordService Password) *Services {
	return &Services{
		Password: passwordService,
	}
}
