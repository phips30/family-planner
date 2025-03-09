package shoppinglist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"family-planner/backend/internal/user"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShoppinglistRouter struct {
	router        *mux.Router
	mongoDbClient *mongo.Database
}

type ShoppinglistItemDto struct {
	Name    string       `json:"name"`
	AddedBy user.UserDto `json:"addedBy"`
	AddedAt time.Time    `json:"addedAt"`
	Bought  bool         `json:"bought"`
	Sorter  int          `json:"sorter"`
}

func NewShoppinglistRouter(router *mux.Router, mongoDbClient *mongo.Database) *ShoppinglistRouter {
	shoppinglistRouter := &ShoppinglistRouter{
		router:        router,
		mongoDbClient: mongoDbClient,
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
		fmt.Println("error parsing request")
		http.Error(w, "error parsing request", http.StatusBadRequest)
	}

	fmt.Println(shoppinglistRequest)
}
