package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NikolayStepanov/PasswordGenerator/internal/domain/dto"
)

type getNewPasswordCommand interface {
	GetNewPassword(ctx context.Context, options dto.GeneratePasswordOptions) (string, error)
}

// GetNewPasswordHandler handles HTTP requests for retrieving passwords
type GetNewPasswordHandler struct {
	name                  string
	getNewPasswordCommand getNewPasswordCommand
}

func (gp *GetNewPasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gp.getNewPassword(w, r)
}

func (gp *GetNewPasswordHandler) getNewPassword(w http.ResponseWriter, r *http.Request) {
	var (
		opts dto.GeneratePasswordOptions
		pass string
	)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&opts)
	if err != nil {
		http.Error(w, "Invalid JSON input: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !opts.IncludeUppercase && !opts.IncludeLowercase && !opts.IncludeDigits {
		http.Error(w, "At least one character type must be included", http.StatusBadRequest)
		return
	}

	pass, err = gp.getNewPasswordCommand.GetNewPassword(r.Context(), opts)
	if err != nil {
		http.Error(w, "Failed to generate password: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"password": pass}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// NewGetPasswordHandler creates and returns a new instance of GetNewPasswordHandler
// with the provided handler name and getNewPasswordCommand implementation
func NewGetPasswordHandler(name string, command getNewPasswordCommand) *GetNewPasswordHandler {
	return &GetNewPasswordHandler{
		name:                  name,
		getNewPasswordCommand: command,
	}
}
