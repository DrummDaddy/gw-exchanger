package storages

type Storage interface {
	GetExchangeRates(fromCurrency, toCurrency string) (float64, error)
	GetAllExchangeRates() (map[string]float64, error)
}
