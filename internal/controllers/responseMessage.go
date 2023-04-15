package controllers

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/iKairat/RecipeAPI/internal/db/sqlc"
)

const (
	msgInvalidParams     = "invalid params"
	msgUserNotFound      = "user not found"
	msgNotFound          = "not found"
	msgIncorrectPswd     = "incorrect password"
	msgUserAlreadyExists = "user already exists"
	msgUnauthenticated   = "unauthenticated"
	msgSuccess           = "success"
)

func wrapMessage(msg string) fiber.Map {
	return fiber.Map{
		"message": msg,
	}
}

type requestBody struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Ingredients []string        `json:"ingredients"`
	Steps       json.RawMessage `json:"steps"`
	TotalTime   int16           `json:"total_time"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type response struct {
	Amount  int         `json:"amount"`
	Status  int         `json:"status"`
	Recipes []db.Recipe `json:"recipes"`
}
