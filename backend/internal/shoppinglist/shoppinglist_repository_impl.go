package shoppinglist

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShoppinglistRepositoryImpl struct {
	ctx               context.Context
	mongoDbCollection *mongo.Collection
}

type shoppinglistMongoItem struct {
	Name    string    `bson:"name"`
	AddedAt time.Time `bson:"addedAt"`
	UserId  string    `bson:"userId"`
	Bought  bool      `bson:"bought"`
}

const MONGO_DB_COLLECTION = "family-planner"

func NewShoppinglistRepositoryImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppinglistRepositoryImpl {
	return &ShoppinglistRepositoryImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistRepositoryImpl) Insert(shoppinglistItems []ShoppinglistItem) ([]ShoppinglistItem, error) {
	mongoShoppingListitems := s.mapToMongoBsonObject(shoppinglistItems)

	result, err := s.mongoDbCollection.InsertMany(s.ctx, mongoShoppingListitems)

	if err != nil {
		return nil, err
	}

	if len(result.InsertedIDs) == len(shoppinglistItems) {
		return shoppinglistItems, nil
	}

	return nil, fmt.Errorf("could not insert all documents into mongodb")
}

func (s *ShoppinglistRepositoryImpl) FindAllForUserIds(userIds []uuid.UUID) ([]ShoppinglistItem, error) {
	var results []shoppinglistMongoItem
	findOptions := options.Find()

	filter := bson.M{"userId": bson.M{"$in": []string{userIds[0].String(), "881b4eb9-d847-4afa-8994-12c6b7307767"}}}
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	for cursor.Next(s.ctx) {
		var item shoppinglistMongoItem
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
	return s.mapToDomainObject(results), nil
}

func (s *ShoppinglistRepositoryImpl) mapToMongoBsonObject(shoppinglistItems []ShoppinglistItem) []interface{} {
	var shoppinglistItemMongoInterfaces []interface{}
	for _, item := range shoppinglistItems {
		shoppinglistMongoItem := shoppinglistMongoItem{
			Name:    item.Name,
			AddedAt: item.AddedAt,
			UserId:  item.UserId.String(),
			Bought:  item.Bought,
		}
		shoppinglistItemMongoInterfaces = append(shoppinglistItemMongoInterfaces, shoppinglistMongoItem)
	}
	return shoppinglistItemMongoInterfaces
}

func (s *ShoppinglistRepositoryImpl) mapToDomainObject(shoppinglistMongoItems []shoppinglistMongoItem) []ShoppinglistItem {
	var shoppinglistItems []ShoppinglistItem
	for _, item := range shoppinglistMongoItems {
		userId, _ := uuid.Parse(item.UserId)
		shoppinglistItem := ShoppinglistItem{
			Name:    item.Name,
			AddedAt: item.AddedAt,
			UserId:  userId,
			Bought:  item.Bought,
		}
		shoppinglistItems = append(shoppinglistItems, shoppinglistItem)
	}
	return shoppinglistItems
}
