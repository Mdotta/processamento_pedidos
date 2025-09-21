package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"processamento_pedidos/internal/models"

	"github.com/google/uuid"
)

func (h Handlers) registerUserEndpoints() {
	http.HandleFunc("GET /users", h.getAllUsers)
	http.HandleFunc("GET /users/query", h.GetUserByID)
	http.HandleFunc("POST /users", h.addUser)
}

func (h Handlers) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.useCases.User.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h Handlers) addUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})

		return
	}

	id, err := h.useCases.User.Add(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateUserResponse{NewUserID: id})
}

func (h Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("invalid user id", "id", idStr, "error", err)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "invalid user id"})
		return
	}
	user, err := h.useCases.User.GetById(userId)
	if err != nil || user == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "user not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.ToGetUserResponse())
}
