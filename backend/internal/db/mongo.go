package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGO_CONNECTION_STRING = "mongodb://localhost:27017/"

func ConnectMongo() (*mongo.Database, error) {
	// Todo: make sure this is only called once and always returns the same dbpool after first init
	clientOptions := options.Client().
		ApplyURI(MONGO_CONNECTION_STRING).
		SetAuth(options.Credential{Username: "mongo", Password: "mongo"})
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo db: %s", err.Error())
	}
	return mongoClient.Database("family-planner"), nil
}
