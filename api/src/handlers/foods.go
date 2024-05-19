package handlers

import (
	database "api/src/database"
	responses "api/src/types"

	"github.com/gofiber/fiber/v2"
)

type food struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
}

func SetupFoodRoutes(router *fiber.App) {
	router.Get("/", GetFoods)
	router.Get("/:id", GetFood)
	router.Post("/", AddFood)
	router.Put("/:id", EditFood)
	router.Delete("/:id", DeleteFood)
}

// Gets all the foods in the database.
func GetFoods(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("SELECT * FROM food;")
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var foods []food = []food{}
	defer rows.Close()
	for rows.Next() {
		var (
			id       int
			name     string
			quantity float32
		)
		if err := rows.Scan(&id, &name, &quantity); err != nil {
			return responses.Failure(ctx, err)
		}
		foods = append(foods, food{Id: id, Name: name, Quantity: quantity})
	}

	return responses.Success(ctx, "Foods retrieved successfully", foods)
}

// Gets a single food from the database.
func GetFood(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("SELECT * FROM food WHERE id = ?;", ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var food food = food{}
	defer rows.Close()
	rows.Next()
	var (
		id       int
		name     string
		quantity float32
	)
	if err := rows.Scan(&id, &name, &quantity); err != nil {
		return responses.Failure(ctx, err)
	}
	food.Id = id
	food.Name = name
	food.Quantity = quantity

	return responses.Success(ctx, "Food retrieved successfully", food)
}

// Adds a new food to the database
func AddFood(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query("INSERT INTO food (name, quantity) VALUES (?, ?) RETURNING *", payload["name"], payload["quantity"])
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var newFood food = food{}
	defer rows.Close()
	rows.Next()
	var (
		id       int
		name     string
		quantity float32
	)
	if err := rows.Scan(&id, &name, &quantity); err != nil {
		return responses.Failure(ctx, err)
	}
	newFood.Id = id
	newFood.Name = name
	newFood.Quantity = quantity

	return responses.Success(ctx, "Food added successfully", newFood)
}

// Edits an existing food in the database
func EditFood(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query("UPDATE food SET name = ?, quantity = ? WHERE id = ? RETURNING *", payload["name"], payload["quantity"], ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var editedFood food = food{}
	defer rows.Close()
	rows.Next()
	var (
		id       int
		name     string
		quantity float32
	)
	if err := rows.Scan(&id, &name, &quantity); err != nil {
		return responses.Failure(ctx, err)
	}
	editedFood.Id = id
	editedFood.Name = name
	editedFood.Quantity = quantity

	return responses.Success(ctx, "Food edited successfully", editedFood)
}

// Deletes a food from the database.
func DeleteFood(ctx *fiber.Ctx) error {
	_, err := database.Connection.Query("DELETE FROM food WHERE id = ?;", ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Food deleted successfully", nil)
}
