package main

import (
	"fmt"
	"net/http"
	"os"

	"family-planner/backend/internal/db"
	"family-planner/backend/internal/user"

	"github.com/gorilla/mux"
)

// To be defined via ENV vars
const PORT = "8080"

func main() {
	dbpool := db.Connect()
	defer dbpool.Close()

	fmt.Printf("Initializing database ...\n")
	err := db.InitDb(dbpool)
	if err != nil {
		fmt.Printf("Error initializing database: %s\n", err)
	}
	fmt.Printf("Initializing database completed.\n")

	// Define Repositories
	userRepository := user.NewUserRepositoryImpl(dbpool)

	// Define Services
	userService := user.NewUserService(userRepository)

	// Define Routing
	r := mux.NewRouter()
	user.NewUserRouter(r, *userService)

	fmt.Printf("Starting server on port %s\n", PORT)
	err = http.ListenAndServe(":"+PORT, r)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
