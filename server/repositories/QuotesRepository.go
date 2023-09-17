package repositories

import (
	"client-server-api/server/models"
	"client-server-api/server/utils"
	"database/sql"
	"github.com/google/uuid"
)

type QuotesRepository struct {
	db *sql.DB
}

func NewQuotesRepository(db *sql.DB) *QuotesRepository {
	return &QuotesRepository{db}
}

func (repository *QuotesRepository) SaveCurrentQuote(quote *models.CurrencyQuote) error {
	stmt, err := repository.db.Prepare("INSERT INTO CurrencyQuotes " +
		"(id, code, codein, name, high, low, var_Bid, pctChange, bid, ask, timestamp, create_Date) " +
		"VALUES " +
		"(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := utils.CreateContext(10)
	defer cancel()

	_, err = stmt.ExecContext(ctx, uuid.New().String(), &quote.Code, &quote.Codein, &quote.Name, &quote.High, &quote.Low, &quote.VarBid, &quote.PctChange, &quote.Bid, &quote.Ask, &quote.Timestamp, &quote.Timestamp)
	if err != nil {
		return err
	}

	return nil
}
