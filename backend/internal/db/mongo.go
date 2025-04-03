package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(connectionString string, dbName string, username string, password string) (*mongo.Database, error) {
	// Todo: make sure this is only called once and always returns the same dbpool after first init
	clientOptions := options.Client().
		ApplyURI(connectionString).
		SetAuth(options.Credential{Username: username, Password: password})
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo db: %s", err.Error())
	}
	return mongoClient.Database(dbName), nil
}
