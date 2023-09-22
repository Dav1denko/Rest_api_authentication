package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	client *mongo.Client
}

func NewAuthMongo(client *mongo.Client) *AuthMongo {
	return &AuthMongo{client: client}
}

func (r *AuthMongo) SaveTokens(GUID int, RefreshToken []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := r.client.Database("admin").Collection("RefreshTokens")
	collection.InsertOne(ctx, bson.D{{Key: "GUID", Value: GUID}, {Key: "RefreshToken", Value: RefreshToken}})
}
