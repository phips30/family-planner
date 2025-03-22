package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain/entity"
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

func (s *ShoppinglistRepositoryImpl) Insert(shoppinglistItems []entity.ShoppinglistItem) error {
	mongoShoppingListitems := MapToMongoBsonObject(shoppinglistItems)
	result, err := s.mongoDbCollection.InsertMany(s.ctx, mongoShoppingListitems)

	if err != nil {
		return err
	}

	if len(result.InsertedIDs) == len(shoppinglistItems) {
		return nil
	}

	return fmt.Errorf("could not insert all documents into mongodb")
}
