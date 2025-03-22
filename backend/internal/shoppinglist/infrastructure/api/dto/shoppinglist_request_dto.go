package dto

import (
	"time"

	"github.com/google/uuid"
)

type ShoppinglistItemRequestDto struct {
	Name    string    `json:"name"`
	UserId  uuid.UUID `json:"userId"`
	GroupId uuid.UUID `json:"groupId"`
	AddedAt time.Time `json:"addedAt"`
	Bought  bool      `json:"bought"`
}
