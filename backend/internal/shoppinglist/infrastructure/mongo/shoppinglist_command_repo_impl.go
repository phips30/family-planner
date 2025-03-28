package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistRepositoryImpl struct {
	ctx               context.Context
	mongoDbCollection *mongo.Collection
}

func NewShoppinglistCommandRepositoryImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppinglistRepositoryImpl {
	return &ShoppinglistRepositoryImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistRepositoryImpl) Insert(shoppinglistItems []entity.ShoppinglistItem) ([]string, error) {
	mongoShoppingListitems := MapToMongoBsonObject(shoppinglistItems)
	result, err := s.mongoDbCollection.InsertMany(s.ctx, mongoShoppingListitems)

	if err != nil {
		return nil, err
	}

	if len(result.InsertedIDs) == len(shoppinglistItems) {
		idsAsString := make([]string, 0, len(result.InsertedIDs))
		for _, id := range result.InsertedIDs {
			idsAsString = append(idsAsString, id.(primitive.ObjectID).Hex())
		}
		return idsAsString, nil
	}

	return nil, fmt.Errorf("could not insert all documents into mongodb")
}

func (s *ShoppinglistRepositoryImpl) Update(shoppingItem entity.ShoppinglistItem) error {
	objectId, err := primitive.ObjectIDFromHex(shoppingItem.Id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"name": shoppingItem.Name, "bought": shoppingItem.Bought}}

	_, err = s.mongoDbCollection.UpdateOne(s.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *ShoppinglistRepositoryImpl) Delete(itemId string) error {
	objectId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	_, err = s.mongoDbCollection.DeleteOne(s.ctx, filter)

	if err != nil {
		return err
	}
	return nil
}
