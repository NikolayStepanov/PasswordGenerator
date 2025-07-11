package handler

import (
	"context"
	"net/http"
)

type getPasswordCommand interface {
	GetPassword(ctx context.Context) (string,error)
}

// GetPasswordHandler handles HTTP requests for retrieving passwords
type GetPasswordHandler struct {
	name               string
	getPasswordCommand getPasswordCommand
}

func (gp *GetPasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gp.getPassword(w,r)
}

func (gp *GetPasswordHandler) getPassword(w http.ResponseWriter, r *http.Request) {

}
// NewGetPasswordHandler creates and returns a new instance of GetPasswordHandler
// with the provided handler name and getPasswordCommand implementation
func NewGetPasswordHandler(name string, command getPasswordCommand) *GetPasswordHandler {
	return &GetPasswordHandler{
		name:               name,
		getPasswordCommand: command,
	}
}
