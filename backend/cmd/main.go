package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"family-planner/backend/internal/db"
	"family-planner/backend/internal/user"

	"github.com/gorilla/mux"
    "github.com/gorilla/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
)

// To be defined via ENV vars
const PORT = "8080"

var (
	dbpool *pgxpool.Pool
	ctx    context.Context = context.Background()
)

func main() {
	dbpool = db.Connect()

	fmt.Printf("Initializing database ...\n")
	err := db.InitDb(dbpool)
	if err != nil {
		fmt.Printf("Error initializing database: %s\n", err)
	} else {
		fmt.Printf("Initializing database completed.\n")
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

	fmt.Printf("Starting server on port %s\n", PORT)
	err = http.ListenAndServe(":"+PORT, handlers.CORS(headers, methods, origins)(r))
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
