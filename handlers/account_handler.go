package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"banking-service/services"
)

// AccountHandler управляет HTTP-запросами, связанными со счетами
type AccountHandler struct {
	AccountService *services.AccountService
}

// CreateAccount обрабатывает создание нового счета
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из контекста (предполагается, что он добавлен middleware)
	userIDStr := r.Context().Value("userID").(string)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	// Создаем счет через сервис
	account, err := h.AccountService.CreateAccount(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем созданный счет в ответе
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// GetUserAccounts обрабатывает получение всех счетов пользователя
func (h *AccountHandler) GetUserAccounts(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из контекста
	userIDStr := r.Context().Value("userID").(string)
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Неверный ID пользователя", http.StatusBadRequest)
		return
	}

	// Получаем счета пользователя через сервис
	accounts, err := h.AccountService.GetUserAccounts(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем список счетов в ответе
	json.NewEncoder(w).Encode(accounts)
}
