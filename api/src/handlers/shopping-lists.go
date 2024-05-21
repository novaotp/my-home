package handlers

import (
	database "api/src/database"
	responses "api/src/types"

	"github.com/gofiber/fiber/v2"
)

type ShoppingList struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type ShoppingListFood struct {
	Id             int     `json:"id"`
	ShoppingListId int     `json:"shoppingListId"`
	FoodId         int     `json:"foodId"`
	Priority       int     `json:"priority"`
	Purchased      bool    `json:"purchased"`
	QuantityToBuy  float32 `json:"quantityToBuy"`
	Notes          string  `json:"notes"`
}

func SetupShoppingListRoutes(router *fiber.App) {
	router.Get("/", GetShoppingLists)
	router.Get("/:id", GetShoppingList)
	router.Post("/", AddShoppingList)
	router.Put("/:id", EditShoppingList)
	router.Delete("/:id", DeleteShoppingList)

	router.Get("/:id/items", GetShoppingListItems)
	router.Post("/:id/items", AddShoppingListItem)
	router.Put("/:id/items/:item_id", EditShoppingListItem)
	router.Delete("/:id/items/:item_id", DeleteShoppingListItem)
}

// Gets all shopping lists from the database.
func GetShoppingLists(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("SELECT * FROM shopping_list;")
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var lists []ShoppingList = []ShoppingList{}
	defer rows.Close()
	for rows.Next() {
		var list ShoppingList
		if err := rows.Scan(&list.Id, &list.Title); err != nil {
			return responses.Failure(ctx, err)
		}
		lists = append(lists, list)
	}

	return responses.Success(ctx, "Shopping lists retrieved successfully", lists)
}

// Gets a single shopping list from the database.
func GetShoppingList(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("SELECT * FROM shopping_list WHERE id = ?;", ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var list ShoppingList = ShoppingList{}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&list.Id, &list.Title); err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Shopping list retrieved successfully", list)
}

// Adds a new shopping list to the database.
func AddShoppingList(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query(
		"INSERT INTO shopping_list (title) VALUES (?) RETURNING *",
		payload["title"],
	)
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var newList ShoppingList = ShoppingList{}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&newList.Id, &newList.Title); err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Shopping list added successfully", newList)
}

// Edits an existing shopping list in the database.
func EditShoppingList(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query(
		"UPDATE shopping_list SET title = ? WHERE id = ? RETURNING *",
		payload["title"],
		ctx.Params("id"),
	)
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var editedList ShoppingList = ShoppingList{}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&editedList.Id, &editedList.Title); err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Shopping list edited successfully", editedList)
}

// Deletes a shopping list from the database.
func DeleteShoppingList(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("DELETE FROM shopping_list WHERE id = ? RETURNING *;", ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var deletedList ShoppingList = ShoppingList{}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&deletedList.Id, &deletedList.Title); err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Shopping list deleted successfully", deletedList)
}

// Gets all items of a shopping list from the database.
func GetShoppingListItems(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("SELECT * FROM shopping_list_food WHERE shopping_list_id = ?;", ctx.Params("id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var items []ShoppingListFood = []ShoppingListFood{}
	defer rows.Close()
	for rows.Next() {
		var item ShoppingListFood
		if err := rows.Scan(&item.Id, &item.ShoppingListId, &item.FoodId, &item.Priority, &item.Purchased, &item.QuantityToBuy, &item.Notes); err != nil {
			return responses.Failure(ctx, err)
		}
		items = append(items, item)
	}

	return responses.Success(ctx, "Shopping list items retrieved successfully", items)
}

// Adds a new item to a shopping list in the database.
func AddShoppingListItem(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query(
		"INSERT INTO shopping_list_food (shopping_list_id, food_id, priority, purchased, quantity_to_buy, notes) VALUES (?, ?, ?, ?, ?, ?) RETURNING *",
		ctx.Params("id"),
		payload["food_id"],
		payload["priority"],
		payload["purchased"],
		payload["quantity_to_buy"],
		payload["notes"],
	)
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var newItem ShoppingListFood = ShoppingListFood{}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&newItem.Id, &newItem.ShoppingListId, &newItem.FoodId, &newItem.Priority, &newItem.Purchased, &newItem.QuantityToBuy, &newItem.Notes); err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Shopping list item added successfully", newItem)
}

// Edits an existing item in a shopping list in the database.
func EditShoppingListItem(ctx *fiber.Ctx) error {
	var payload map[string]any
	if err := ctx.BodyParser(&payload); err != nil {
		return responses.Failure(ctx, err)
	}

	rows, err := database.Connection.Query(
		"UPDATE shopping_list_food SET food_id = ?, priority = ?, purchased = ?, quantity_to_buy = ?, notes = ? WHERE shopping_list_id = ? AND food_id = ? RETURNING *",
		payload["food_id"],
		payload["priority"],
		payload["purchased"],
		payload["quantity_to_buy"],
		payload["notes"],
		ctx.Params("id"),
		ctx.Params("item_id"),
	)
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var editedItem ShoppingListFood = ShoppingListFood{}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&editedItem.Id, &editedItem.ShoppingListId, &editedItem.FoodId, &editedItem.Priority, &editedItem.Purchased, &editedItem.QuantityToBuy, &editedItem.Notes); err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Shopping list item edited successfully", editedItem)
}

// Deletes an item from a shopping list in the database.
func DeleteShoppingListItem(ctx *fiber.Ctx) error {
	rows, err := database.Connection.Query("DELETE FROM shopping_list_food WHERE shopping_list_id = ? AND food_id = ? RETURNING *;", ctx.Params("id"), ctx.Params("item_id"))
	if err != nil {
		return responses.Failure(ctx, err)
	}

	var deletedItem ShoppingListFood = ShoppingListFood{}
	defer rows.Close()
	rows.Next()
	if err := rows.Scan(&deletedItem.Id, &deletedItem.ShoppingListId, &deletedItem.FoodId, &deletedItem.Priority, &deletedItem.Purchased, &deletedItem.QuantityToBuy, &deletedItem.Notes); err != nil {
		return responses.Failure(ctx, err)
	}

	return responses.Success(ctx, "Shopping list item deleted successfully", deletedItem)
}
