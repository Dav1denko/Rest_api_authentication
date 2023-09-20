package service

import "Rest_api_authentication/pkg/repository"

type Authorization interface{}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
