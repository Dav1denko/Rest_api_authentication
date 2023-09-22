package repository

import "go.mongodb.org/mongo-driver/mongo"

type AuthMongo struct {
	client *mongo.Client
}

func NewAuthMongo(client *mongo.Client) *AuthMongo {
	return &AuthMongo{client: client}
}

func (r *AuthMongo) GetTokens() {
	print("test")
}
