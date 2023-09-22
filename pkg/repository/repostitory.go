package repository

import "go.mongodb.org/mongo-driver/mongo"

type Authorization interface {
	SaveTokens(GUID int, RefreshToken []byte)
}

type Repository struct {
	Authorization
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(client),
	}
}
