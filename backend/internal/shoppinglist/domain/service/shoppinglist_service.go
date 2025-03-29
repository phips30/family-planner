package service

import (
	"errors"
	"family-planner/backend/internal/common/models"
	"family-planner/backend/internal/shoppinglist/domain"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"family-planner/backend/internal/shoppinglist/domain/repository"
	"log"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistService struct {
	commandRepository  repository.ShoppinglistCommandRepository
	queryRepository    repository.ShoppinglistQueryRepository
	userDataRepository repository.UserDataRepository
}

var (
	errGettingItem            = errors.New("error getting shopping list item")
	errRetrievingListForUser  = errors.New("error getting shopping list for user")
	errRetrievingListForGroup = errors.New("error getting shopping list for group")
	errGettingUserData        = errors.New("error loading user data for shopping list items")
	errNoItemIdProvided       = errors.New("error no item id provided")
)

func NewShoppinglistService(commandRepository repository.ShoppinglistCommandRepository,
	queryRepository repository.ShoppinglistQueryRepository,
	userDataRepository repository.UserDataRepository) *ShoppinglistService {
	return &ShoppinglistService{
		commandRepository:  commandRepository,
		queryRepository:    queryRepository,
		userDataRepository: userDataRepository,
	}
}

func (s *ShoppinglistService) LoadShoppingListForUserId(userid uuid.UUID) ([]domain.ShoppinglistItemResponse, error) {
	items, err := s.queryRepository.FindAllByUserId(userid)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingListForUser
	}

	var userIds []uuid.UUID
	for _, item := range items {
		userIds = append(userIds, item.UserId.UUID)
	}
	uniqueUserIds := s.getUniqueUserIds(userIds)
	itemCreators, err := s.userDataRepository.GetUserData(uniqueUserIds)

	var shoppinglistItems []domain.ShoppinglistItemResponse
	var itemCreator models.UserDto
	for _, item := range items {
		for _, user := range itemCreators {
			if user.Id == item.UserId.UUID {
				itemCreator = models.UserDto{Id: user.Id, Name: user.Name, Email: user.Email}
				break
			}
		}
		responseItem := &domain.ShoppinglistItemResponse{
			Id:      item.Id,
			Name:    string(item.Name),
			User:    itemCreator,
			AddedAt: item.AddedAt,
			Bought:  item.Bought,
		}

		shoppinglistItems = append(shoppinglistItems, *responseItem)

	}
	return shoppinglistItems, nil
}

func (s *ShoppinglistService) LoadShoppingListForGroupId(groupId uuid.UUID) ([]domain.ShoppinglistItemResponse, error) {
	items, err := s.queryRepository.FindAllByGroupId(groupId)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingListForGroup
	}

	var userIds []uuid.UUID
	for _, item := range items {
		userIds = append(userIds, item.UserId.UUID)
	}
	uniqueUserIds := s.getUniqueUserIds(userIds)
	itemCreators, err := s.userDataRepository.GetUserData(uniqueUserIds)

	var shoppinglistItems []domain.ShoppinglistItemResponse
	var itemCreator models.UserDto
	for _, item := range items {
		for _, user := range itemCreators {
			if user.Id == item.UserId.UUID {
				itemCreator = models.UserDto{Id: user.Id, Name: user.Name, Email: user.Email}
				break
			}
		}
		responseItem := &domain.ShoppinglistItemResponse{
			Id:      item.Id,
			Name:    string(item.Name),
			User:    itemCreator,
			AddedAt: item.AddedAt,
			Bought:  item.Bought,
		}

		shoppinglistItems = append(shoppinglistItems, *responseItem)

	}
	return shoppinglistItems, nil
}

func (s *ShoppinglistService) getUniqueUserIds(allUserIds []uuid.UUID) []uuid.UUID {
	seenUserIds := make(map[uuid.UUID]bool)
	var uniqueUserIds []uuid.UUID
	for _, userId := range allUserIds {
		if !seenUserIds[userId] {
			seenUserIds[userId] = true
			uniqueUserIds = append(uniqueUserIds, userId)
		}
	}
	return uniqueUserIds
}

func (s *ShoppinglistService) SaveShoppingList(shoppinglist []domain.ShoppinglistDto) ([]string, error) {
	var items []entity.ShoppinglistItem
	for _, item := range shoppinglist {
		newItem, err := entity.NewShoppinglistItem(
			item.Name,
			item.AddedAt,
			item.UserId,
			item.GroupId,
			item.Bought,
		)
		if err != nil {
			continue
		}

		items = append(items, *newItem)
	}
	return s.commandRepository.Insert(items)
}

func (s *ShoppinglistService) UpdateItem(itemId string, name string, bought bool) error {
	item, err := s.queryRepository.FindById(itemId)
	if err != nil || item == nil || item.Id == "" {
		return errGettingItem
	}

	itemName, err := entity.NewItemName(name)
	if err != nil {
		return err
	}

	item.SetName(itemName)
	item.SetBought(bought)
	return s.commandRepository.Update(*item)
}

func (s *ShoppinglistService) DeleteItem(itemId string) error {
	if itemId == "" {
		return errNoItemIdProvided
	}
	return s.commandRepository.Delete(itemId)
}

func (s *ShoppinglistService) validate(shoppinglistItems []domain.ShoppinglistDto) error {
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
