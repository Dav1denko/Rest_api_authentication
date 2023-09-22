package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	collection.InsertOne(ctx, bson.D{{Key: "GUID", Value: GUID}, {Key: "RefreshToken", Value: string(RefreshToken)}})
}

func (r *AuthMongo) GetRefreshTokens(GUID int, CookieRefreshToken string) bool {
	var result struct {
		GUID         int    `bson:"GUID"`
		RefreshToken string `bson:"RefreshToken"`
	}
	filter := bson.D{{Key: "GUID", Value: GUID}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := r.client.Database("admin").Collection("RefreshTokens")
	collection.FindOne(ctx, filter).Decode(&result)
	verification := bcrypt.CompareHashAndPassword([]byte(result.RefreshToken), []byte(CookieRefreshToken))
	if verification == nil {
		collection.DeleteOne(ctx, bson.D{{Key: "GUID", Value: GUID}})
		return true
	}
	return false

}
