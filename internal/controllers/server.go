package controllers

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/iKairat/RecipeAPI/internal/db/sqlc"
	"github.com/iKairat/RecipeAPI/internal/middleware"
)

type Server struct {
	App   *fiber.App
	store db.Store
}

func New(dbconn db.Store, app *fiber.App) *Server {

	s := &Server{
		store: dbconn,
	}
	app.Post("/register", s.register)
	app.Post("/login", s.login)

	// checks for authenticated users
	app.Use(middleware.IsAuthenticated)

	app.Post("/recipe", s.createRecipe)
	app.Get("/recipes/id/:id", s.getRecipe)
	app.Get("/recipes/all", s.getAllRecipes)
	app.Get("/recipes/", s.getRecipesByIngredients)
	app.Get("/recipes/time/", s.getRecipesByTotaltime)
	app.Put("/recipe/:id", s.updateRecipe)
	app.Delete("/recipe/:id", s.deleteRecipe)

	app.Post("/logout", s.logout)
	app.Get("/users", s.getUsers)

	s.App = app

	return s
}
