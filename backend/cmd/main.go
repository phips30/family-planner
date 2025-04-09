package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"family-planner/backend/internal/common/config"
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

var (
	dbpool *pgxpool.Pool
	ctx    context.Context = context.Background()
)

func main() {
	// Load config vars
	config.Init("./internal/common/config/.env")
	cfg := config.GetConfig()

	// Init postgres
	dbpool, err := db.ConnectPostgres(cfg.PostgresConnectionString)
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
	mongoDbClient, err := db.ConnectMongo(cfg.MongoConnectionString, cfg.MongoDatabase, cfg.MongoUsername, cfg.MongoPassword)
	if err != nil {
		log.Fatal("Error connecting to mongo", err.Error())
	} else {
		log.Println("Connected to MongoDB!")
	}

	// Define Repositories
	userRepository := postgres.NewUserRepositoryImpl(ctx, dbpool)
	groupRepository := postgres.NewGroupRepositoryImpl(ctx, dbpool, userRepository)
	userDataAdapter := postgres.NewUserDataAdapter(ctx, dbpool)
	shoppinglistCommandRepo := mongo.NewShoppinglistCommandRepositoryImpl(ctx, mongoDbClient)
	shoppinglistQueryRepo := mongo.NewShoppinglistQueryRepoImpl(ctx, mongoDbClient)

	// Define Services
	userService := familyService.NewUserService(userRepository)
	groupService := familyService.NewGroupService(groupRepository, *userService)
	shoppinglistService := shoppinglistService.NewShoppinglistService(
		shoppinglistCommandRepo,
		shoppinglistQueryRepo,
		userDataAdapter,
	)

	// Define Routing
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	familyApi.NewUserRouter(r, *userService, *groupService)
	familyApi.NewGroupRouter(r, *groupService)
	shoppinglistApi.NewShoppinglistRouter(r, mongoDbClient, shoppinglistService)

	log.Printf("Starting server on port %s", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, handlers.CORS(headers, methods, origins)(r))
	if err != nil {
		log.Printf("Error starting server: %s", err)
		os.Exit(1)
	}
}
