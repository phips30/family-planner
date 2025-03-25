package service

import (
	"errors"
	"family-planner/backend/internal/shoppinglist/domain"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"family-planner/backend/internal/shoppinglist/domain/port"
	"family-planner/backend/internal/shoppinglist/domain/repository"
	"log"
	"time"

	"github.com/google/uuid"
)

type ShoppinglistService struct {
	commandRepository repository.ShoppinglistCommandRepository
	queryRepository   repository.ShoppinglistQueryRepository
	userDataPort      port.UserDataPort
}

var (
	errGettingItem            = errors.New("error getting shopping list item")
	errRetrievingListForUser  = errors.New("error getting shopping list for user")
	errRetrievingListForGroup = errors.New("error getting shopping list for group")
	errGettingUserData        = errors.New("error loading user data for shopping list items")
	errNoItemIdProvided       = errors.New("error no item id provided")
)

type ShoppinglistItemCreator struct {
	Id    uuid.UUID
	Name  string
	Email string
}

type ShoppinglistItemResponse struct {
	Id      string
	Name    string
	AddedAt time.Time
	User    ShoppinglistItemCreator
	Bought  bool
}

func NewShoppinglistService(commandRepository repository.ShoppinglistCommandRepository,
	queryRepository repository.ShoppinglistQueryRepository,
	userDataPort port.UserDataPort) *ShoppinglistService {
	return &ShoppinglistService{
		commandRepository: commandRepository,
		queryRepository:   queryRepository,
		userDataPort:      userDataPort,
	}
}

func (s *ShoppinglistService) LoadShoppingListForUserId(userid uuid.UUID) ([]ShoppinglistItemResponse, error) {
	items, err := s.queryRepository.FindAllByUserId(userid)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingListForUser
	}

	// This can be done in the repo already
	domainItems := entity.FromExistingItems(items)

	var userIds []uuid.UUID
	for _, item := range items {
		userIds = append(userIds, item.UserId)
	}
	uniqueUserIds := s.getUniqueUserIds(userIds)
	itemCreators, err := s.userDataPort.GetUserData(uniqueUserIds)

	var shoppinglistItems []ShoppinglistItemResponse
	var itemCreator ShoppinglistItemCreator
	for _, item := range *domainItems {
		for _, user := range itemCreators {
			if user.Id.UUID == item.UserId.UUID {
				itemCreator = ShoppinglistItemCreator{Id: user.Id.UUID, Name: user.Name, Email: user.Email}
				break
			}
		}
		responseItem := &ShoppinglistItemResponse{
			Id:      item.Id,
			Name:    item.Name.ToString(),
			User:    itemCreator,
			AddedAt: item.AddedAt,
			Bought:  item.Bought,
		}

		shoppinglistItems = append(shoppinglistItems, *responseItem)

	}
	return shoppinglistItems, nil
}

func (s *ShoppinglistService) LoadShoppingListForGroupId(groupId uuid.UUID) ([]ShoppinglistItemResponse, error) {
	items, err := s.queryRepository.FindAllByGroupId(groupId)
	if err != nil {
		log.Println(err.Error())
		return nil, errRetrievingListForGroup
	}

	// This can be done in the repo already
	domainItems := entity.FromExistingItems(items)

	var userIds []uuid.UUID
	for _, item := range items {
		userIds = append(userIds, item.UserId)
	}
	uniqueUserIds := s.getUniqueUserIds(userIds)
	itemCreators, err := s.userDataPort.GetUserData(uniqueUserIds)

	var shoppinglistItems []ShoppinglistItemResponse
	var itemCreator ShoppinglistItemCreator
	for _, item := range *domainItems {
		for _, user := range itemCreators {
			if user.Id.UUID == item.UserId.UUID {
				itemCreator = ShoppinglistItemCreator{Id: user.Id.UUID, Name: user.Name, Email: user.Email}
				break
			}
		}
		responseItem := &ShoppinglistItemResponse{
			Id:      item.Id,
			Name:    item.Name.ToString(),
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

func (s *ShoppinglistService) SaveShoppingList(shoppinglist []domain.ShoppinglistDto) error {
	var items []entity.ShoppinglistItem
	for _, item := range shoppinglist {
		name, err := entity.NewItemName(item.Name)
		if err != nil {
			break
		}

		userId, err := entity.NewUserId(item.UserId)
		if err != nil {
			break
		}

		newItem, err := entity.NewShoppinglistItem(
			*name,
			item.AddedAt,
			*userId,
			*entity.NewGroupId(item.GroupId),
			item.Bought,
		)

		items = append(items, *newItem)

	}
	return s.commandRepository.Insert(items)
}

func (s *ShoppinglistService) UpdateItem(itemId string, name string, bought bool) error {
	item, err := s.queryRepository.FindById(itemId)
	if err != nil || item == nil {
		return errGettingItem
	}
	shoppingItem := entity.FromExistingItem(*item)
	itemName, err := entity.NewItemName(name)
	if err != nil {
		return err
	}

	shoppingItem.SetName(*itemName)
	shoppingItem.SetBought(bought)
	return s.commandRepository.Update(*shoppingItem)
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
