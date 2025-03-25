package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistMongoItem struct {
	Id      string    `bson:"_id"`
	Name    string    `bson:"name"`
	AddedAt time.Time `bson:"addedAt"`
	UserId  string    `bson:"userId"`
	GroupId string    `bson:"groupId"`
	Bought  bool      `bson:"bought"`
}

func ExtractCursorIntoDomainObject(cursor *mongo.Cursor, ctx context.Context) ([]domain.ShoppinglistDto, error) {
	var shoppinglistDtos []domain.ShoppinglistDto
	for cursor.Next(ctx) {
		var item ShoppinglistMongoItem
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		shoppinglistDtos = append(shoppinglistDtos, MapToDto(item))
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return shoppinglistDtos, nil
}

func MapToMongoBsonObject(shoppinglistItems []entity.ShoppinglistItem) []interface{} {
	var shoppinglistItemMongoInterfaces []interface{}
	for _, item := range shoppinglistItems {
		shoppinglistMongoItem := ShoppinglistMongoItem{
			Name:    item.Name.ToString(),
			AddedAt: item.AddedAt,
			UserId:  item.UserId.String(),
			GroupId: item.GroupId.String(),
			Bought:  item.Bought,
		}
		shoppinglistItemMongoInterfaces = append(shoppinglistItemMongoInterfaces, shoppinglistMongoItem)
	}
	return shoppinglistItemMongoInterfaces
}

func MapToDto(shoppinglistMongoItem ShoppinglistMongoItem) domain.ShoppinglistDto {
	userId, _ := uuid.Parse(shoppinglistMongoItem.UserId)
	groupId, _ := uuid.Parse(shoppinglistMongoItem.GroupId)
	return domain.ShoppinglistDto{
		Id:      shoppinglistMongoItem.Id,
		Name:    shoppinglistMongoItem.Name,
		AddedAt: shoppinglistMongoItem.AddedAt,
		UserId:  userId,
		GroupId: groupId,
		Bought:  shoppinglistMongoItem.Bought,
	}
}
