package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"family-planner/backend/internal/db"
	familyService "family-planner/backend/internal/family/domain/service"
	familyApi "family-planner/backend/internal/family/infrastructure/api"
	"family-planner/backend/internal/family/infrastructure/postgres"
	shoppinglistService "family-planner/backend/internal/shoppinglist/domain/service"
	shoppinglistApi "family-planner/backend/internal/shoppinglist/infrastructure/api"
	"family-planner/backend/internal/shoppinglist/infrastructure/mongo"

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
	dbpool, err := db.ConnectPostgres()
	if err != nil {
		log.Fatal("Error connecting to postgres", err.Error())
	} else {
		log.Println("Connected to Postgres!")
	}

	log.Println("Initializing database ...")
	if err = db.InitDb(dbpool); err != nil {
		log.Printf("Error initializing database: %s\n", err)
	} else {
		log.Println("Initializing database completed.")
	}

	// Init mongodb
	mongoDbClient, err := db.ConnectMongo()
	if err != nil {
		log.Fatal("Error connecting to mongo", err.Error())
	} else {
		log.Println("Connected to MongoDB!")
	}

	// Define Repositories
	userRepository := postgres.NewUserRepositoryImpl(ctx, dbpool)
	groupRepository := postgres.NewGroupRepositoryImpl(ctx, dbpool, userRepository)
	shoppinglistSaveRepo := mongo.NewShoppinglistRepositoryImpl(ctx, mongoDbClient)
	shoppinglistQueryUserRepo := mongo.NewShoppinglistQueryUserRepositoryImpl(ctx, mongoDbClient)
	shoppinglistQueryGroupRepo := mongo.NewShoppinglistQueryGroupRepoImpl(ctx, mongoDbClient)

	// Define Services
	userService := familyService.NewUserService(userRepository)
	groupService := familyService.NewGroupService(groupRepository, *userService)
	shoppinglistService := shoppinglistService.NewShoppinglistService(
		shoppinglistSaveRepo,
		shoppinglistQueryUserRepo,
		shoppinglistQueryGroupRepo)

	// Define Routing
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	familyApi.NewUserRouter(r, *userService)
	familyApi.NewGroupRouter(r, *groupService)
	shoppinglistApi.NewShoppinglistRouter(r, mongoDbClient, shoppinglistService)

	log.Printf("Starting server on port %s", PORT)
	err = http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origins)(r))
	if err != nil {
		log.Printf("Error starting server: %s", err)
		os.Exit(1)
	}
}
