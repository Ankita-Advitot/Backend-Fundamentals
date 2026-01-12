package main

import (
	"net/http"
	"modules/controllers"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", controllers.HealthCheck)
	mux.HandleFunc("/users", controllers.UsersHandler)
	mux.HandleFunc("/users/", controllers.UserByIDHandler) // route param

	return mux
}