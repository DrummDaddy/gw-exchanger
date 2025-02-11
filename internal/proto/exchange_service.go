package proto

import (
	"context"
	"gw-exchanger/internal/storages"

	pb "github.com/DrummDaddy/proto-exchange/exchange"
)

type ExchangeService struct {
	pb.UnimplementedExchangeServiceServer
	storage storages.Storage
}

func NewExchangeService(storage storages.Storage) *ExchangeService {
	return &ExchangeService{storage: storage}
}

func (s *ExchangeService) GetExchangeRates(ctx context.Context, _ *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates, err := s.storage.GetAllExchangeRates()
	if err != nil {
		return nil, err
	}

	response := &pb.ExchangeRatesResponse{Rates: make(map[string]float32)}
	for key, val := range rates {
		response.Rates[key] = float32(val)
	}
	return response, nil

}

func (s *ExchangeService) GetExchangeRateForCurrency(ctx context.Context, req *pb.CurrencyRequest) (*pb.ExchangeRateResponse, error) {
	rate, err := s.storage.GetExchangeRates(req.FromCurrency, req.ToCurrency)
	if err != nil {
		return nil, err
	}

	return &pb.ExchangeRateResponse{
		FromCurrency: req.FromCurrency,
		ToCurrency:   req.ToCurrency,
		Rate:         float32(rate),
	}, nil
}
