package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/common/dto"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistMongoItem struct {
	Name    string    `bson:"name"`
	AddedAt time.Time `bson:"addedAt"`
	UserId  string    `bson:"userId"`
	GroupId string    `bson:"groupId"`
	Bought  bool      `bson:"bought"`
}

func ExtractCursorIntoDomainObject(cursor *mongo.Cursor, ctx context.Context) ([]dto.ShoppinglistItemRequestDto, error) {
	var shoppinglistItems []dto.ShoppinglistItemRequestDto
	for cursor.Next(ctx) {
		var item ShoppinglistMongoItem
		err := cursor.Decode(&item)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		shoppinglistItems = append(shoppinglistItems, MapToDomainObject(item))
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return shoppinglistItems, nil
}

func MapToMongoBsonObject(shoppinglistItems []entity.ShoppinglistItem) []interface{} {
	var shoppinglistItemMongoInterfaces []interface{}
	for _, item := range shoppinglistItems {
		shoppinglistMongoItem := ShoppinglistMongoItem{
			Name:    item.Name,
			AddedAt: item.AddedAt,
			UserId:  item.User.Id.String(),
			GroupId: item.GroupId.String(),
			Bought:  item.Bought,
		}
		shoppinglistItemMongoInterfaces = append(shoppinglistItemMongoInterfaces, shoppinglistMongoItem)
	}
	return shoppinglistItemMongoInterfaces
}

func MapToDomainObject(shoppinglistMongoItem ShoppinglistMongoItem) dto.ShoppinglistItemRequestDto {
	userId, _ := uuid.Parse(shoppinglistMongoItem.UserId)
	groupId, _ := uuid.Parse(shoppinglistMongoItem.GroupId)
	return dto.ShoppinglistItemRequestDto{
		Name:    shoppinglistMongoItem.Name,
		AddedAt: shoppinglistMongoItem.AddedAt,
		UserId:  userId,
		GroupId: groupId,
		Bought:  shoppinglistMongoItem.Bought,
	}
}
