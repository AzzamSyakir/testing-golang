package routes

import (
	"database/sql"
	"net/http"
	"testing-golang/application/controller"
	"testing-golang/application/middleware"
	"testing-golang/application/repositories"
	"testing-golang/application/service"
	"testing-golang/config"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {
	// Initialize repositories
	userRepository := repositories.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(*userRepository)

	// Initialize controllers
	userController := controller.NewUserController(*userService)

	// Create a new router
	router := mux.NewRouter()

	// Protected routes
	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)

	// Authentication routes
	router.HandleFunc("/users", userController.CreateUserController).Methods("POST")
	router.HandleFunc("/users/login", userController.LoginUser).Methods("POST")
	router.HandleFunc("/users/logout", userController.LogoutUser).Methods("POST")

	// User routes
	protectedRoutes.HandleFunc("/users", userController.FetchUserController).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", userController.GetUserController).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", userController.UpdateUserController).Methods("PUT")
	protectedRoutes.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}

func RunServer() {
	db := config.InitDB()
	router := Router(db)

	// Mulai server HTTP dengan router yang telah dikonfigurasi
	http.Handle("/", router)
	http.ListenAndServe(":9000", nil)
}
