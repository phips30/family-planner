package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
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

func NewShoppinglistRepositoryImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppinglistRepositoryImpl {
	return &ShoppinglistRepositoryImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistRepositoryImpl) Insert(shoppinglistItems []entity.ShoppinglistItem) ([]entity.ShoppinglistItem, error) {
	mongoShoppingListitems := MapToMongoBsonObject(shoppinglistItems)

	result, err := s.mongoDbCollection.InsertMany(s.ctx, mongoShoppingListitems)

	if err != nil {
		return nil, err
	}

	if len(result.InsertedIDs) == len(shoppinglistItems) {
		return shoppinglistItems, nil
	}

	return nil, fmt.Errorf("could not insert all documents into mongodb")
}
