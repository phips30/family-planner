package mongo

import (
	"context"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistMongoItem struct {
	Id      string    `bson:"_id,omitempty"`
	Name    string    `bson:"name"`
	AddedAt time.Time `bson:"addedAt"`
	UserId  string    `bson:"userId"`
	GroupId string    `bson:"groupId"`
	Bought  bool      `bson:"bought"`
}

func ExtractCursorIntoDomainObject(cursor *mongo.Cursor, ctx context.Context) ([]entity.ShoppinglistItem, error) {
	var shoppinglistItems []entity.ShoppinglistItem
	for cursor.Next(ctx) {
		var item ShoppinglistMongoItem
		err := cursor.Decode(&item)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		shoppinglistItemDO, err := MapToDomainObject(item)
		if err != nil {
			continue
		} else {
			shoppinglistItems = append(shoppinglistItems, *shoppinglistItemDO)
		}
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

func MapToDomainObject(mongoItem ShoppinglistMongoItem) (*entity.ShoppinglistItem, error) {
	userId, _ := uuid.Parse(mongoItem.UserId)
	groupId, _ := uuid.Parse(mongoItem.GroupId)
	return entity.FromExistingItem(
		mongoItem.Id,
		mongoItem.Name,
		mongoItem.AddedAt,
		userId,
		groupId,
		mongoItem.Bought,
	)
}
