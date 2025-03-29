package entity

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserId struct {
	uuid.UUID
}

func (u UserId) ToString() string {
	return u.String()
}

func NewUserId(id uuid.UUID) (UserId, error) {
	if id == (uuid.UUID{}) {
		return UserId{}, errors.New("no user id provided")
	}
	userId := UserId{id}
	return userId, nil
}

func NewUserIdFromString(id string) (UserId, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return UserId{}, fmt.Errorf("could not parse uuid: %s", id)
	}
	return NewUserId(uuid)
}
