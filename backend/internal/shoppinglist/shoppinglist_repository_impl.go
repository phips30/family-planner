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
	GroupId string    `bson:"groupId"`
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

func (s *ShoppinglistRepositoryImpl) FindAllForUserIds(userId uuid.UUID) ([]ShoppinglistItem, error) {
	findOptions := options.Find()
	filter := bson.M{"userId": userId.String()}
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	return s.extractCursorIntoDomainObject(cursor)
}

func (s *ShoppinglistRepositoryImpl) FindAllForGroupIds(groupId uuid.UUID) ([]ShoppinglistItem, error) {
	findOptions := options.Find()
	filter := bson.M{"groupId": groupId.String()}
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	return s.extractCursorIntoDomainObject(cursor)
}

func (s *ShoppinglistRepositoryImpl) extractCursorIntoDomainObject(cursor *mongo.Cursor) ([]ShoppinglistItem, error) {
	var shoppinglistItems []ShoppinglistItem
	for cursor.Next(s.ctx) {
		var item shoppinglistMongoItem
		err := cursor.Decode(&item)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		shoppinglistItems = append(shoppinglistItems, s.mapToDomainObject(item))
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return shoppinglistItems, nil
}

func (s *ShoppinglistRepositoryImpl) mapToMongoBsonObject(shoppinglistItems []ShoppinglistItem) []interface{} {
	var shoppinglistItemMongoInterfaces []interface{}
	for _, item := range shoppinglistItems {
		shoppinglistMongoItem := shoppinglistMongoItem{
			Name:    item.Name,
			AddedAt: item.AddedAt,
			UserId:  item.UserId.String(),
			GroupId: item.GroupId.String(),
			Bought:  item.Bought,
		}
		shoppinglistItemMongoInterfaces = append(shoppinglistItemMongoInterfaces, shoppinglistMongoItem)
	}
	return shoppinglistItemMongoInterfaces
}

func (s *ShoppinglistRepositoryImpl) mapToDomainObject(shoppinglistMongoItem shoppinglistMongoItem) ShoppinglistItem {
	userId, _ := uuid.Parse(shoppinglistMongoItem.UserId)
	groupId, _ := uuid.Parse(shoppinglistMongoItem.GroupId)
	return ShoppinglistItem{
		Name:    shoppinglistMongoItem.Name,
		AddedAt: shoppinglistMongoItem.AddedAt,
		UserId:  userId,
		GroupId: groupId,
		Bought:  shoppinglistMongoItem.Bought,
	}
}
