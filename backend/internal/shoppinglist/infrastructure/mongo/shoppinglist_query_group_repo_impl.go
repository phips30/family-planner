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

type ShoppinglistQueryGroupRepoImpl struct {
	ctx               context.Context
	mongoDbCollection *mongo.Collection
}

func NewShoppinglistQueryGroupRepoImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppinglistQueryGroupRepoImpl {
	return &ShoppinglistQueryGroupRepoImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistQueryGroupRepoImpl) FindAll(groupId uuid.UUID) ([]domain.ShoppinglistDto, error) {
	findOptions := options.Find()
	filter := bson.M{"groupId": groupId.String()}
	cursor, err := s.mongoDbCollection.Find(s.ctx, filter, findOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(s.ctx)

	return ExtractCursorIntoDomainObject(cursor, s.ctx)
}
