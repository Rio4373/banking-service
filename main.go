package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"banking-service/config"
	"banking-service/handlers"
	"banking-service/middleware"
	"banking-service/repositories"
	"banking-service/services"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация подключения к базе данных
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	// Инициализация репозиториев
	transactionRepo := &repositories.TransactionRepository{DB: db}

	// Инициализация сервисов
	transferService := &services.TransferService{TransactionRepo: transactionRepo}

	// Инициализация обработчиков
	transferHandler := &handlers.TransferHandler{TransferService: transferService}

	// Настройка маршрутизатора
	r := mux.NewRouter()

	// Защищенные эндпоинты (требуют JWT-аутентификацию)
	authRouter := r.PathPrefix("/").Subrouter()
	authRouter.Use(middleware.AuthMiddleware)
	authRouter.HandleFunc("/transfer", transferHandler.Transfer).Methods("POST")

	// Запуск сервера
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
