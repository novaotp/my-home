package handlers

import (
	database "api/src/database"
	responses "api/src/types"

	"github.com/gofiber/fiber/v2"
)

type Food struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Unit     string  `json:"unit"`
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

	var foods []Food = []Food{}
	defer rows.Close()
	for rows.Next() {
		var (
			id       int
			name     string
			quantity float32
			unit     string
		)
		if err := rows.Scan(&id, &name, &quantity, &unit); err != nil {
			return responses.Failure(ctx, err)
		}
		foods = append(foods, Food{Id: id, Name: name, Quantity: quantity, Unit: unit})
	}

	return responses.Success(ctx, "Foods retrieved successfully", foods)
}

// Gets a single food from the database.
func GetFood(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("SELECT * FROM food WHERE id = ?;", ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var food Food = Food{}
	defer rows.Close()
	rows.Next()
	var (
		id       int
		name     string
		quantity float32
		unit     string
	)
	if err := rows.Scan(&id, &name, &quantity, &unit); err != nil {
		return responses.Failure(ctx, err)
	}
	food.Id = id
	food.Name = name
	food.Quantity = quantity
	food.Unit = unit

	return responses.Success(ctx, "Food retrieved successfully", food)
}

// Adds a new food to the database
func AddFood(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query(
		"INSERT INTO food (name, quantity, unit) VALUES (?, ?, ?) RETURNING *",
		payload["name"],
		payload["quantity"],
		payload["unit"],
	)
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var newFood Food = Food{}
	defer rows.Close()
	rows.Next()
	var (
		id       int
		name     string
		quantity float32
		unit     string
	)
	if err := rows.Scan(&id, &name, &quantity, &unit); err != nil {
		return responses.Failure(ctx, err)
	}
	newFood.Id = id
	newFood.Name = name
	newFood.Quantity = quantity
	newFood.Unit = unit

	return responses.Success(ctx, "Food added successfully", newFood)
}

// Edits an existing food in the database
func EditFood(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query(
		"UPDATE food SET name = ?, quantity = ?, unit = ? WHERE id = ? RETURNING *",
		payload["name"],
		payload["quantity"],
		payload["unit"],
		ctx.Params("id"),
	)
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var editedFood Food = Food{}
	defer rows.Close()
	rows.Next()
	var (
		id       int
		name     string
		quantity float32
		unit     string
	)
	if err := rows.Scan(&id, &name, &quantity, &unit); err != nil {
		return responses.Failure(ctx, err)
	}
	editedFood.Id = id
	editedFood.Name = name
	editedFood.Quantity = quantity
	editedFood.Unit = unit

	return responses.Success(ctx, "Food edited successfully", editedFood)
}

// Deletes a food from the database.
func DeleteFood(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("DELETE FROM food WHERE id = ? RETURNING *;", ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var deletedFood Food = Food{}
	defer rows.Close()
	rows.Next()
	var (
		id       int
		name     string
		quantity float32
		unit     string
	)
	if err := rows.Scan(&id, &name, &quantity, &unit); err != nil {
		return responses.Failure(ctx, err)
	}
	deletedFood.Id = id
	deletedFood.Name = name
	deletedFood.Quantity = quantity
	deletedFood.Unit = unit

	return responses.Success(ctx, "Food deleted successfully", deletedFood)
}
