package service

import (
	"errors"
	"family-planner/backend/internal/shoppinglist/common/dto"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"family-planner/backend/internal/shoppinglist/domain/port"
	"family-planner/backend/internal/shoppinglist/domain/repository"
	"log"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistService struct {
	saveRepository       repository.ShoppinglistSaveRepository
	queryUserRepository  repository.ShoppinglistQueryRepository
	queryGroupRepository repository.ShoppinglistQueryRepository
	userDataPort         port.UserDataPort
}

var (
	errRetrievingListForUser  = errors.New("error getting shopping list for user")
	errRetrievingListForGroup = errors.New("error getting shopping list for group")
	errGettingUserData        = errors.New("error loading user data for shopping list items")
)

func NewShoppinglistService(saveRepository repository.ShoppinglistSaveRepository,
	queryUserRepository repository.ShoppinglistQueryRepository,
	queryGroupRepository repository.ShoppinglistQueryRepository,
	userDataPort port.UserDataPort) *ShoppinglistService {
	return &ShoppinglistService{
		saveRepository:       saveRepository,
		queryUserRepository:  queryUserRepository,
		queryGroupRepository: queryGroupRepository,
		userDataPort:         userDataPort,
	}
}

func (s *ShoppinglistService) LoadShoppingListForUserId(userid uuid.UUID) ([]entity.ShoppinglistItem, error) {
	items, err := s.queryUserRepository.FindAll(userid)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingListForUser
	}

	userIds := s.getUniqueUserIds(items)
	itemCreators, err := s.userDataPort.GetUserData(userIds)
	return *entity.FromExistingItems(items, itemCreators), nil
}

func (s *ShoppinglistService) LoadShoppingListForGroupId(groupId uuid.UUID) ([]entity.ShoppinglistItem, error) {
	items, err := s.queryGroupRepository.FindAll(groupId)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingListForGroup
	}

	userIds := s.getUniqueUserIds(items)
	itemCreators, err := s.userDataPort.GetUserData(userIds)
	return *entity.FromExistingItems(items, itemCreators), nil
}

func (s *ShoppinglistService) getUniqueUserIds(items []dto.ShoppinglistItemRequestDto) []uuid.UUID {
	var userIds []uuid.UUID
	for _, item := range items {
		userIds = append(userIds, item.UserId)
	}

	seenUserIds := make(map[uuid.UUID]bool)
	var uniqueUserIds []uuid.UUID
	for _, userId := range userIds {
		if !seenUserIds[userId] {
			seenUserIds[userId] = true
			uniqueUserIds = append(uniqueUserIds, userId)
		}
	}
	return userIds
}

func (s *ShoppinglistService) SaveShoppingList(shoppinglist []dto.ShoppinglistItemRequestDto) ([]entity.ShoppinglistItem, error) {
	userIds := s.getUniqueUserIds(shoppinglist)
	itemCreators, err := s.userDataPort.GetUserData(userIds)
	if err != nil {
		return nil, errGettingUserData
	}

	items := entity.NewShoppinglistItems(shoppinglist, itemCreators)

	return s.saveRepository.Insert(*items)
	/*
	   validation_error := s.validate(shoppinglist)
	   // TODO: Just aggregate the errors and log them instead of not continuing

	   	if validation_error != nil {
	   		return nil, fmt.Errorf("%s", validation_error.Error())
	   	}

	   return s.saveRepository.Insert(shoppinglist)
	*/
}

func (s *ShoppinglistService) validate(shoppinglistItems []dto.ShoppinglistItemRequestDto) error {
	// TODO: Check if item does not already exist for the group for a shopping trip
	for _, shoppinglistItem := range shoppinglistItems {
		if shoppinglistItem.Name == "" {
			return errors.New("no name provided")
		}
		if shoppinglistItem.UserId == (uuid.UUID{}) {
			return errors.New("no user provided")
		}
		shoppinglistItem.AddedAt = time.Now()
	}
	return nil
}
