package services

import (
	"client-server-api/server/models"
	"client-server-api/server/repositories"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type QuotesService struct {
	repositories *repositories.QuotesRepository
}

func NewQuotesService(db *sql.DB) *QuotesService {
	repo := repositories.NewQuotesRepository(db)
	return &QuotesService{repo}
}

func (service *QuotesService) GetCurrentQuote(currency string) (map[string]models.CurrencyQuote, error) {
	resp, err := requestQuote(currency)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Printf("Erro ao fechar o corpo da resposta: %v\n", closeErr)
		}
	}()

	var data map[string]models.CurrencyQuote
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	for _, quote := range data {
		err = service.repositories.SaveCurrentQuote(&quote)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func requestQuote(currency string) (*http.Response, error) {

	URL := "https://economia.awesomeapi.com.br/json/last/" + currency
	httpClient := &http.Client{Timeout: 200 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a solicitação: %v", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a solicitação: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta não esperada: %s", resp.Status)
	}

	return resp, nil
}
