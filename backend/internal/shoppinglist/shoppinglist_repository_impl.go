package shoppinglist

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShoppinglistRepositoryImpl struct {
	ctx               context.Context
	mongoDbCollection *mongo.Collection
}

const MONGO_DB_COLLECTION = "family-planner"

func NewShoppinglistRepositoryImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppinglistRepositoryImpl {
	return &ShoppinglistRepositoryImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistRepositoryImpl) Insert(shoppinglistItems []ShoppinglistItem) ([]ShoppinglistItem, error) {
	var interfaces []interface{}
	for _, item := range shoppinglistItems {
		interfaces = append(interfaces, item)
	}

	result, err := s.mongoDbCollection.InsertMany(s.ctx, interfaces)

	if err != nil {
		return nil, err
	}

	if len(result.InsertedIDs) == len(shoppinglistItems) {
		return shoppinglistItems, nil
	}

	return nil, fmt.Errorf("could not insert all documents into mongodb")
}

func (s *ShoppinglistRepositoryImpl) FindAllForUserIds(userIds []uuid.UUID) ([]ShoppinglistItem, error) {
	var results []ShoppinglistItem
	findOptions := options.Find()
	filter := bson.M{"userId": bson.M{"$in": userIds}}
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var item ShoppinglistItem
		err := cursor.Decode(&item)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		results = append(results, item)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	for _, result := range results {
		fmt.Printf("Name: %s, addedby: %s\n", result.Name, result.UserId)
	}
	return results, nil
}
