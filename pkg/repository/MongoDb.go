package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
}

func NewMongoDB(ctx context.Context, cfg Config) (*mongo.Client, error) {
	clientOption := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	clientOption.SetAuth(options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	})

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
