package main

import (
	database "api/src/database"
	handlers "api/src/handlers"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	database.Connect()

	http.HandleFunc("GET /foods", handlers.GetFoods)
	http.HandleFunc("GET /foods/", handlers.GetFood)
	http.HandleFunc("POST /foods", handlers.AddFood)
	http.HandleFunc("PUT /foods/", handlers.EditFood)
	http.HandleFunc("DELETE /foods/", handlers.DeleteFood)

	fmt.Println("> Running on port 8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
