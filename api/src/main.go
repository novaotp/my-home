package main

import (
	database "api/src/database"
	handlers "api/src/handlers"
	"api/src/middlewares"
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "modernc.org/sqlite"
)

func main() {
	app := fiber.New()

	middlewares.HandleCORS(app)

	database.Connect()

	foodsRouter := fiber.New()
	handlers.SetupFoodRoutes(foodsRouter)
	app.Mount("/api/v1/foods", foodsRouter)

	shoppingListsRouter := fiber.New()
	handlers.SetupShoppingListRoutes(shoppingListsRouter)
	app.Mount("/api/v1/shopping-lists", shoppingListsRouter)

	fmt.Println("> Running on port 8080")
	app.Listen("127.0.0.1:8080")
}
