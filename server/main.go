package main

import (
	"client-server-api/server/api/routes"
	_ "client-server-api/server/models"
	"client-server-api/server/utils"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	db, err := sql.Open("sqlite3", "Quotes.db")
	if err != nil {
		log.Fatal("Erro ao abrir o banco de dados:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Erro ao fechar a conexão com o banco de dados: %v", err)
		}
	}(db)

	const PORT = 8080
	utils.PrintBanner()
	fmt.Printf("\nServidor API rodando na porta %d...\n", PORT)

	route := mux.NewRouter()
	routes.SetupRoutes(db, route)

	http.Handle("/", route)
	err = http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		panic(fmt.Sprintf("Erro ao iniciar o servidor: %v", err))
	}

}
