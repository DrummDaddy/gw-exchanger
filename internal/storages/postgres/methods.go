package postgres

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) GetAllExchangeRates() (map[string]float64, error) {
	query := "SELECT from_currency, to_currency, rate FROM currency_rates"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rates := make(map[string]float64)
	for rows.Next() {
		var from, to string
		var rate float64
		if err := rows.Scan(&from, &to, rate); err != nil {
			return nil, err
		}
		rates[from+to] = rate
	}
	return rates, nil

}

func (s *Storage) GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	query := "SELECT rate FROM currency_rates WHERE from_currency = $1 AND to_currency = $2"
	var rate float64
	if err := s.db.QueryRow(query, fromCurrency, toCurrency).Scan(&rate); err != nil {
		return 0, err

	}
	return rate, nil
}

func (s *Storage) GetExchangeRates(fromCurrency, toCurrency string) (float64, error) {
	return s.GetExchangeRate(fromCurrency, toCurrency)
}
