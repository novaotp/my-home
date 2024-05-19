package handlers

import (
	database "api/src/database"
	responses "api/src/types"
	"encoding/json"
	"net/http"
)

type food struct {
	Id       int
	Name     string
	Quantity float32
}

// Gets all the foods in the database.
func GetFoods(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Connection.Query("SELECT * FROM food;")
	if err != nil {
		responses.FailureResponse(w, err)
		return
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
			responses.FailureResponse(w, err)
			return
		}
		foods = append(foods, food{Id: id, Name: name, Quantity: quantity})
	}

	responses.SuccessResponse(w, "Foods retrieved successfully", foods)
}

// Gets a single food from the database.
func GetFood(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Connection.Query("SELECT * FROM food WHERE id = ?;", r.URL.Path[7:])
	if err != nil {
		responses.FailureResponse(w, err)
		return
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
		responses.FailureResponse(w, err)
		return
	}
	food.Id = id
	food.Name = name
	food.Quantity = quantity

	responses.SuccessResponse(w, "Food retrieved successfully", food)
}

// Adds a new food to the database
func AddFood(w http.ResponseWriter, r *http.Request) {
	var data map[string]any
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		responses.FailureResponse(w, err)
		return
	}

	rows, err := database.Connection.Query("INSERT INTO food (name, quantity) VALUES (?, ?) RETURNING *", data["name"], data["quantity"])
	if err != nil {
		responses.FailureResponse(w, err)
		return
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
		responses.FailureResponse(w, err)
		return
	}
	newFood.Id = id
	newFood.Name = name
	newFood.Quantity = quantity

	responses.SuccessResponse(w, "Food added successfully", newFood)
}

// Edits an existing food in the database
func EditFood(w http.ResponseWriter, r *http.Request) {
	var data map[string]any
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		responses.FailureResponse(w, err)
		return
	}

	rows, err := database.Connection.Query("UPDATE food SET name = ?, quantity = ? WHERE id = ? RETURNING *", data["name"], data["quantity"], r.URL.Path[7:])
	if err != nil {
		responses.FailureResponse(w, err)
		return
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
		responses.FailureResponse(w, err)
		return
	}
	editedFood.Id = id
	editedFood.Name = name
	editedFood.Quantity = quantity

	responses.SuccessResponse(w, "Food edited successfully", editedFood)
}

// Deletes a food from the database.
func DeleteFood(w http.ResponseWriter, r *http.Request) {
	_, err := database.Connection.Query("DELETE FROM food WHERE id = ?;", r.URL.Path[7:])
	if err != nil {
		responses.FailureResponse(w, err)
		return
	}

	responses.SuccessResponse(w, "Food deleted successfully", nil)
}
