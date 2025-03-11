package shoppinglist

import (
	"encoding/json"
	"family-planner/backend/internal/user"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistRouter struct {
	router        *mux.Router
	mongoDbClient *mongo.Database
	service       *ShoppinglistService
}

type ShoppinglistItemDto struct {
	Name string `json:"name"`
	// TODO: This can be just the email
	AddedBy user.UserDto `json:"addedBy"`
	AddedAt time.Time    `json:"addedAt"`
	Bought  bool         `json:"bought"`
}

func NewShoppinglistRouter(router *mux.Router, mongoDbClient *mongo.Database, shoppinglistService *ShoppinglistService) *ShoppinglistRouter {
	shoppinglistRouter := &ShoppinglistRouter{
		router:        router,
		mongoDbClient: mongoDbClient,
		service:       shoppinglistService,
	}

	router.HandleFunc("/shopping-list", shoppinglistRouter.findList).Methods("GET")
	router.HandleFunc("/shopping-list", shoppinglistRouter.createList).Methods("POST")

	return shoppinglistRouter
}

func (s *ShoppinglistRouter) findList(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	email := params.Get("email")

	fmt.Printf("Trying to find shopping list - Email: %s\n", email)
}

func (s *ShoppinglistRouter) createList(w http.ResponseWriter, r *http.Request) {
	var shoppinglistRequest []ShoppinglistItemDto

	if err := json.NewDecoder(r.Body).Decode(&shoppinglistRequest); err != nil {
		log.Println("error parsing request")
		http.Error(w, "error parsing request", http.StatusBadRequest)
	}

	log.Println("Wer are ")
	for _, item := range shoppinglistRequest {
		log.Println(item.AddedBy)
	}

	_, err := s.service.SaveShoppingList(shoppinglistRequest)
	if err != nil {
		log.Println("Error saving shopping list:", err.Error())
	}
	json.NewEncoder(w).Encode(shoppinglistRequest)
}
