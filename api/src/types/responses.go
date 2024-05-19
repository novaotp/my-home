package types

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Logs the error and defines a 500 failure response in the writer.
func FailureResponse(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(map[string]any{
		"success": false,
		"message": "Internal Server Error",
	})
}

// Defines a 200 success response in the writer with custom message and optional data.
func SuccessResponse(w http.ResponseWriter, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if data != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"message": message,
			"data":    data,
		})
	} else {
		json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"message": message,
		})
	}
}
