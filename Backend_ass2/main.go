package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := SetupRoutes()

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", router))
}
