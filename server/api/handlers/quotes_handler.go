package handlers

import (
	"client-server-api/server/services"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type QuoteHandler struct {
	service *services.QuotesService
}

func NewQuotesHandler(db *sql.DB) *QuoteHandler {
	service := services.NewQuotesService(db)
	return &QuoteHandler{service}
}

func (handler *QuoteHandler) GetCurrencyQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	resp, err := handler.service.GetCurrentQuote(vars["currency"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao solicitar cotação: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
