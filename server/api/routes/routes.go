package routes

import (
	"client-server-api/server/api/handlers"
	"database/sql"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB, route *mux.Router) {
	quotesHandler := handlers.NewQuotesHandler(db)
	route.HandleFunc("/currencyQuotes/{currency}", quotesHandler.GetCurrencyQuote).Methods("GET")
}
