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

type ShoppinglistItemUpdateRequestDto struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bought bool   `json:"bought"`
}
