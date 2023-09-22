package service

import (
	"Rest_api_authentication/pkg/repository"
	"strconv"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{}
}

func (s *AuthService) GetTokens(GUID int) (string, error) {
	GUID_test := strconv.Itoa(GUID)
	return GUID_test, nil
}
