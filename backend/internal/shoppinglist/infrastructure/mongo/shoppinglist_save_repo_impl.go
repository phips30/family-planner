package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/common/dto"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistRepositoryImpl struct {
	ctx               context.Context
	mongoDbCollection *mongo.Collection
}

func NewShoppinglistRepositoryImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppinglistRepositoryImpl {
	return &ShoppinglistRepositoryImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistRepositoryImpl) Insert(shoppinglistItems []dto.ShoppinglistItemRequestDto) ([]dto.ShoppinglistItemRequestDto, error) {
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

func (s *ShoppinglistRepositoryImpl) mapToMongoBsonObject(shoppinglistItems []dto.ShoppinglistItemRequestDto) []interface{} {
	var shoppinglistItemMongoInterfaces []interface{}
	for _, item := range shoppinglistItems {
		shoppinglistItemMongoInterfaces = append(shoppinglistItemMongoInterfaces, item)
	}
	return shoppinglistItemMongoInterfaces
}
