package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShoppinglistQueryRepoImpl struct {
	ctx               context.Context
	mongoDbCollection *mongo.Collection
}

func NewShoppinglistQueryRepoImpl(ctx context.Context, mongoDbClient *mongo.Database) *ShoppinglistQueryRepoImpl {
	return &ShoppinglistQueryRepoImpl{ctx: ctx, mongoDbCollection: mongoDbClient.Collection(MONGO_DB_COLLECTION)}
}

func (s *ShoppinglistQueryRepoImpl) FindAllByGroupId(groupId uuid.UUID) ([]domain.ShoppinglistDto, error) {
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

func (s *ShoppinglistQueryRepoImpl) FindAllByUserId(userId uuid.UUID) ([]domain.ShoppinglistDto, error) {
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

func (s *ShoppinglistQueryRepoImpl) FindById(itemId string) (*domain.ShoppinglistDto, error) {
	objectId, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return nil, err
	}

	var item ShoppinglistMongoItem
	filter := bson.M{"_id": objectId}
	err = s.mongoDbCollection.FindOne(s.ctx, filter).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	itemDto := MapToDto(item)
	return &itemDto, nil
}
