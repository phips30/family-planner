package shoppinglist

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShoppinglistRepositoryImpl struct {
	ctx               context.Context
	dbpool            *pgxpool.Pool
	mongoDbCollection *mongo.Collection
}

const MONGO_DB_COLLECTION = "family-planner"

func NewShoppinglistRepositoryImpl(ctx context.Context, dbpool *pgxpool.Pool, mongoDbClient *mongo.Database) *ShoppinglistRepositoryImpl {
	return &ShoppinglistRepositoryImpl{ctx: ctx, dbpool: dbpool, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistRepositoryImpl) Create(shoppinglistItems []*ShoppinglistItem) ([]*ShoppinglistItem, error) {
	_, err := s.mongoDbCollection.InsertOne(s.ctx, shoppinglistItems)

	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return shoppinglistItems, nil
}

func (s *ShoppinglistRepositoryImpl) FindAllInGroupForEmail(email string) ([]ShoppinglistItem, error) {
	// Todo: Select all emails in group

	var results []ShoppinglistItem
	findOptions := options.Find()
	filter := bson.M{} // Add filter for email addresses
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	// Iterate through the results
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
	// Print the results
	for _, result := range results {
		fmt.Printf("Name: %s, addedby: %s\n", result.Name, result.AddedBy)
	}
	return results, nil
}
