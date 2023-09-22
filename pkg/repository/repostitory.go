package repository

import "go.mongodb.org/mongo-driver/mongo"

type Authorization interface {
	GetTokens()
}

type Repository struct {
	Authorization
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(client),
	}
}
