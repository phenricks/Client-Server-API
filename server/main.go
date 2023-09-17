package main

import (
	_ "client-server-api/server/models"
	"client-server-api/server/utils"
	"fmt"
	"net/http"
)

func main() {
	const PORT = 8080
	utils.PrintBanner()
	fmt.Printf("\nServidor API rodando na porta %d...\n", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}

}
