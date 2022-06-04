package usecase

import (
	"net/http"

	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/helper"
	"travel-agent-backend/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type serviceAuth struct {
	c    config.Config
	repo adapter.AdapterAuthRepository
}

func (s *serviceAuth) RegisterService(user model.User) (int, error) {
	_, errUsername := s.repo.UsernameExists(user.Username)
	if errUsername != nil {
		return http.StatusExpectationFailed, errUsername
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPassword)

	errRegister := s.repo.Register(user)
	if errRegister != nil {
		return http.StatusInternalServerError, errRegister
	}

	return http.StatusOK, nil
}

func (s *serviceAuth) LoginService(username string, password string) (string, int) {
	user, err := s.repo.Login(username)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(user.ID, user.Username, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func NewServiceAuth(repo adapter.AdapterAuthRepository, c config.Config) adapter.AdapterAuthService {
	return &serviceAuth{
		repo: repo,
		c:    c,
	}
}
