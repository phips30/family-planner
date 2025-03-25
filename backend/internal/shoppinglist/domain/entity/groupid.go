package entity

import (
	"fmt"

	"github.com/google/uuid"
)

type GroupId struct {
	uuid.UUID
}

func (g GroupId) ToString() string {
	return g.String()
}

func NewGroupId(id uuid.UUID) *GroupId {
	groupId := GroupId{id}
	return &groupId
}

func NewGroupIdFromString(id string) (*GroupId, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("could not parse uuid: %s", id)
	}
	return NewGroupId(uuid), nil
}
