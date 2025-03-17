package service

import (
	"errors"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"family-planner/backend/internal/shoppinglist/domain/repository"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistService struct {
	saveRepository       repository.ShoppinglistSaveRepository
	queryUserRepository  repository.ShoppinglistQueryRepository
	queryGroupRepository repository.ShoppinglistQueryRepository
}

var (
	errRetrievingUserData  = errors.New("error getting shopping list for user")
	errRetrievingGroupData = errors.New("error getting shopping list for group")
)

func NewShoppinglistService(saveRepository repository.ShoppinglistSaveRepository,
	queryUserRepository repository.ShoppinglistQueryRepository,
	queryGroupRepository repository.ShoppinglistQueryRepository) *ShoppinglistService {
	return &ShoppinglistService{
		saveRepository:       saveRepository,
		queryUserRepository:  queryUserRepository,
		queryGroupRepository: queryGroupRepository,
	}
}

func (s *ShoppinglistService) LoadShoppingListForUserId(userid uuid.UUID) ([]entity.ShoppinglistItem, error) {
	items, err := s.queryUserRepository.FindAll(userid)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingUserData
	}
	return items, nil
}

func (s *ShoppinglistService) LoadShoppingListForGroupId(groupId uuid.UUID) ([]entity.ShoppinglistItem, error) {
	items, err := s.queryGroupRepository.FindAll(groupId)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingGroupData
	}
	return items, nil
}

func (s *ShoppinglistService) SaveShoppingList(shoppinglist []entity.ShoppinglistItem) ([]entity.ShoppinglistItem, error) {
	validation_error := s.validate(shoppinglist)
	// TODO: Just aggregate the errors and log them instead of not continuing
	if validation_error != nil {
		return nil, fmt.Errorf("%s", validation_error.Error())
	}
	return s.saveRepository.Insert(shoppinglist)
}

func (s *ShoppinglistService) validate(shoppinglistItems []entity.ShoppinglistItem) error {
	// TODO: Check if item does not already exist for the group for a shopping trip
	for _, shoppinglistItem := range shoppinglistItems {
		if shoppinglistItem.Name == "" {
			return errors.New("no name provided")
		}
		if shoppinglistItem.UserId == (uuid.UUID{}) {
			return errors.New("no user provided")
		}
		if shoppinglistItem.AddedAt == (time.Time{}) {
			shoppinglistItem.AddedAt = time.Now()
		}
	}
	return nil
}
