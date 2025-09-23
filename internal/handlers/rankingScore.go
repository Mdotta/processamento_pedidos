package handlers

import (
	"encoding/json"
	"net/http"
	"processamento_pedidos/internal/models"
)

func (h Handlers) registerrankingScoreEndpoints() {
	http.HandleFunc("GET /ranking", h.getRanking)
	http.HandleFunc("POST /ranking", h.addRanking)
}

func (h Handlers) getRanking(w http.ResponseWriter, r *http.Request) {
	rankings := h.useCases.RankingScore.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rankings)
}

func (h Handlers) addRanking(w http.ResponseWriter, r *http.Request) {
	var req models.CreateRankingScoreRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})

		return
	}

	id, err := h.useCases.RankingScore.Add(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.CreateRankingScoreResponse{RankingId: id})
}
