package handlers

import (
	database "api/src/database"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type food struct {
	Id       int
	Name     string
	Quantity float32
}

// Gets all the foods in the database.
func GetFoods(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := db.Query("SELECT * FROM food;")
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Internal Server Error",
		})
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
			panic(err)
		}
		foods = append(foods, food{Id: id, Name: name, Quantity: quantity})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Foods retrieved successfully",
		"data":    foods,
	})
}

// Gets a single food from the database.
func GetFood(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows, err := db.Query("SELECT * FROM food WHERE id = ?;", r.URL.Path[7:])
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Internal Server Error",
		})
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
		panic(err)
	}
	food.Id = id
	food.Name = name
	food.Quantity = quantity

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Food retrieved successfully",
		"data":    food,
	})
}

// Adds a new food to the database
func AddFood(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data map[string]any
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Internal Server Error",
		})
		return
	}

	rows, err := db.Query("INSERT INTO food (name, quantity) VALUES (?, ?) RETURNING *", data["name"], data["quantity"])
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Internal Server Error",
		})
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
		panic(err)
	}
	newFood.Id = id
	newFood.Name = name
	newFood.Quantity = quantity

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Food added successfully",
		"data":    newFood,
	})
}

// Edits an existing food in the database
func EditFood(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data map[string]any
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Internal Server Error",
		})
		return
	}

	rows, err := db.Query("UPDATE food SET name = ?, quantity = ? WHERE id = ? RETURNING *", data["name"], data["quantity"], r.URL.Path[7:])
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Internal Server Error",
		})
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
		panic(err)
	}
	editedFood.Id = id
	editedFood.Name = name
	editedFood.Quantity = quantity

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Food edited successfully",
		"data":    editedFood,
	})
}

// Deletes a food from the database.
func DeleteFood(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Query("DELETE FROM food WHERE id = ?;", r.URL.Path[7:])
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(map[string]any{
			"success": false,
			"message": "Internal Server Error",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{
		"success": true,
		"message": "Food deleted successfully",
	})
}
