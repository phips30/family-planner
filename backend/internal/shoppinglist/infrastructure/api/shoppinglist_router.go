package api

import (
	"encoding/json"
	"family-planner/backend/internal/shoppinglist/domain"
	"family-planner/backend/internal/shoppinglist/domain/entity"
	"family-planner/backend/internal/shoppinglist/domain/service"
	"family-planner/backend/internal/shoppinglist/infrastructure/api/dto"

	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistRouter struct {
	router        *mux.Router
	mongoDbClient *mongo.Database
	service       *service.ShoppinglistService
}

func NewShoppinglistRouter(router *mux.Router, mongoDbClient *mongo.Database, shoppinglistService *service.ShoppinglistService) *ShoppinglistRouter {
	shoppinglistRouter := &ShoppinglistRouter{
		router:        router,
		mongoDbClient: mongoDbClient,
		service:       shoppinglistService,
	}

	router.HandleFunc("/shopping-list/user/{userId}", shoppinglistRouter.findListForUser).Methods("GET")
	router.HandleFunc("/shopping-list/group/{groupId}", shoppinglistRouter.findListForGroup).Methods("GET")
	router.HandleFunc("/shopping-list/", shoppinglistRouter.addItems).Methods("POST")
	//router.HandleFunc("/shopping-list/{}/{id}", shoppinglistRouter.updateItems).Methods("PUT")
	router.HandleFunc("/shopping-list/{itemId}", shoppinglistRouter.deleteItem).Methods("DELETE")

	return shoppinglistRouter
}

func (s *ShoppinglistRouter) findListForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := uuid.Parse(vars["userId"])
	if err != nil {
		log.Println("No proper uuid provided: ", vars["userId"])
		http.Error(w, "No proper uuid provided", http.StatusBadRequest)
	}

	shoppinglistItems, err := s.service.LoadShoppingListForUserId(userId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var shoppinglistResponse []dto.ShoppinglistItemResponseDto
	for _, shoppinlistItem := range shoppinglistItems {
		shoppinglistResponse = append(shoppinglistResponse, *s.mapFromDomainObject(shoppinlistItem))
	}

	json.NewEncoder(w).Encode(shoppinglistResponse)
}

func (s *ShoppinglistRouter) findListForGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupId, err := uuid.Parse(vars["groupId"])
	if err != nil {
		log.Println("No proper uuid provided: ", vars["groupId"])
		http.Error(w, "No proper uuid provided", http.StatusBadRequest)
	}

	shoppinglistItems, err := s.service.LoadShoppingListForGroupId(groupId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var shoppinglistResponse []dto.ShoppinglistItemResponseDto
	for _, shoppinlistItem := range shoppinglistItems {
		shoppinglistResponse = append(shoppinglistResponse, *s.mapFromDomainObject(shoppinlistItem))
	}

	json.NewEncoder(w).Encode(shoppinglistResponse)
}

func (s *ShoppinglistRouter) addItems(w http.ResponseWriter, r *http.Request) {
	var shoppinglistRequest []dto.ShoppinglistItemRequestDto
	if err := json.NewDecoder(r.Body).Decode(&shoppinglistRequest); err != nil {
		log.Println("error parsing request")
		http.Error(w, "error parsing request", http.StatusBadRequest)
	}

	err := s.service.SaveShoppingList(s.mapToDomainObject(shoppinglistRequest))
	if err != nil {
		log.Println("Error saving shopping list:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func (s *ShoppinglistRouter) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]

	err := s.service.DeleteItem(itemId)
	if err != nil {
		log.Println("Error deleting item:", err.Error())
		w.WriteHeader(http.StatusBadRequest)

	}
	w.WriteHeader(http.StatusOK)
}

func (s *ShoppinglistRouter) mapFromDomainObject(shoppinglistItem entity.ShoppinglistItem) *dto.ShoppinglistItemResponseDto {
	return &dto.ShoppinglistItemResponseDto{
		Id:      shoppinglistItem.Id,
		Name:    shoppinglistItem.Name,
		User:    dto.ShoppingtItemCreatorResponseDto{Id: shoppinglistItem.User.Id, Name: shoppinglistItem.User.Name, Email: shoppinglistItem.User.Email},
		AddedAt: shoppinglistItem.AddedAt,
		Bought:  shoppinglistItem.Bought,
	}
}

func (s *ShoppinglistRouter) mapToDomainObject(shoppinglistItems []dto.ShoppinglistItemRequestDto) []domain.ShoppinglistDto {
	var domainShoppingListItems []domain.ShoppinglistDto
	for _, item := range shoppinglistItems {
		domainShoppingListItem := domain.ShoppinglistDto{
			Name:    item.Name,
			UserId:  item.UserId,
			GroupId: item.GroupId,
			AddedAt: item.AddedAt,
			Bought:  item.Bought,
		}
		domainShoppingListItems = append(domainShoppingListItems, domainShoppingListItem)
	}
	return domainShoppingListItems
}
