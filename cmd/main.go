package main

import (
	"log"
	"net"

	pb "github.com/DrummDaddy/proto-exchange/exchange"
	"google.golang.org/grpc"

	"gw-exchanger/internal/config"
	"gw-exchanger/internal/proto"
	"gw-exchanger/internal/storages/postgres"
)

func main() {
	// Чтение и парсинг конфигурации
	cfg, err := config.LoadConfig("config.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// Инициализация подключения к PostgreSQL
	db, err := postgres.NewConnector(cfg.DB)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	storage := postgres.NewStorage(db)

	// Инициализируем gRPC сервер
	server := grpc.NewServer()
	exchangeService := proto.NewExchangeService(storage)
	pb.RegisterExchangeServiceServer(server, exchangeService)

	// Запуск сервера
	listener, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Ошибка запуска gRPC сервера: %v", err)
	}
	log.Printf("Сервер запущен на %s", cfg.GRPCPort)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("gRPC Server ошибка: %v", err)
	}
}
