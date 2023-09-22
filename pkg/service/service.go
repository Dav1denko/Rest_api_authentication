package service

import (
	restapiauthentication "Rest_api_authentication"
	"Rest_api_authentication/pkg/repository"
)

type Authorization interface {
	GetTokens(GUID int) (restapiauthentication.Info, error)
	SaveRefreshToken(GUID int, RefreshToken string)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
