package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShoppingListQueryUserRepoImpl struct {
	ctx               context.Context
	mongoDbCollection *mongo.Collection
}

func NewShoppinglistQueryUserRepositoryImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppingListQueryUserRepoImpl {
	return &ShoppingListQueryUserRepoImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppingListQueryUserRepoImpl) FindAll(id uuid.UUID) ([]entity.ShoppinglistItem, error) {
	findOptions := options.Find()
	filter := bson.M{"userId": id.String()}
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	return ExtractCursorIntoDomainObject(cursor, s.ctx)
}
