package service

import "Rest_api_authentication/pkg/repository"

type Authorization interface {
	GetTokens(GUID int) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
