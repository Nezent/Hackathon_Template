package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Nezent/Hackathon_Template/internal/application/dto"
	"github.com/Nezent/Hackathon_Template/internal/domain/user"
	"github.com/Nezent/Hackathon_Template/pkg/response"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Implement handler methods here
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.service.CreateUser(&req)
	if err != nil {
		response.WriteError(w, err.Error(), err.StatusCode)
		return
	}

	response.WriteSuccess(w, res, http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetUser()
	if err != nil {
		response.WriteError(w, err.Error(), err.StatusCode)
		return
	}
	response.WriteSuccess(w, res, http.StatusOK)
}
