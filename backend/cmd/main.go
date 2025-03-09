package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"family-planner/backend/internal/db"
	"family-planner/backend/internal/shoppinglist"
	"family-planner/backend/internal/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// To be defined via ENV vars
const PORT = "8080"

var (
	dbpool *pgxpool.Pool
	ctx    context.Context = context.Background()
)

func main() {
	// Init postgres
	dbpool = db.ConnectPostgres()
	log.Println("Initializing database ...")
	err := db.InitDb(dbpool)
	if err != nil {
		log.Printf("Error initializing database: %s\n", err)
	} else {
		log.Println("Initializing database completed.")
	}

	// Init mongodb
	mongoDbClient, err := db.ConnectMongo()
	if err != nil {
		log.Panic("Error connecting to mongo\n", err.Error())
	} else {
		log.Println("Connected to MongoDB!")
	}

	// Define Repositories
	userRepository := user.NewUserRepositoryImpl(ctx, dbpool)

	// Define Services
	userService := user.NewUserService(userRepository)

	// Define Routing
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	user.NewUserRouter(r, *userService)
	shoppinglist.NewShoppinglistRouter(r, mongoDbClient)

	fmt.Printf("Starting server on port %s\n", PORT)
	err = http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origins)(r))
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
