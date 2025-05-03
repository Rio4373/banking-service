package services

import (
	"banking-service/models"
	"banking-service/repositories"
	"errors"
)

// AuthService управляет процессами регистрации и аутентификации
type AuthService struct {
	UserRepo *repositories.UserRepository
}

// Register регистрирует нового пользователя
func (s *AuthService) Register(username, email, password string) error {
	user := &models.User{
		Username: username,
		Email:    email,
	}
	if err := user.HashPassword(password); err != nil {
		return errors.New("не удалось захешировать пароль")
	}

	existingUser, _ := s.UserRepo.GetUserByEmail(email)
	if existingUser != nil {
		return errors.New("email уже используется")
	}

	return s.UserRepo.CreateUser(user)
}

// Login аутентифицирует пользователя и возвращает JWT-токен
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("неверные данные для входа")
	}

	if !user.CheckPassword(password) {
		return "", errors.New("неверные данные для входа")
	}

	// Генерация JWT-токена (реализация опущена для краткости)
	return "jwt_token", nil
}
