package types

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Logs the error and defines a 500 failure response.
func Failure(ctx *fiber.Ctx, err error) error {
	fmt.Println(err)
	return ctx.Status(500).JSON(map[string]any{
		"success": false,
		"message": "Internal Server Error",
	})
}

// Defines a 200 success response, sending a custom message and data (optional).
func Success(ctx *fiber.Ctx, message string, data any) error {
	var response map[string]any = map[string]any{
		"success": true,
		"message": message,
	}
	if data != nil {
		response["data"] = data
	}
	return ctx.Status(200).JSON(response)
}
