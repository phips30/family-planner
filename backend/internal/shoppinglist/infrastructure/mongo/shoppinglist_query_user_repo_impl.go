package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain"
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

func (s *ShoppingListQueryUserRepoImpl) FindAll(userId uuid.UUID) ([]domain.ShoppinglistDto, error) {
	findOptions := options.Find()
	filter := bson.M{"userId": userId.String()}
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	return ExtractCursorIntoDomainObject(cursor, s.ctx)
}
