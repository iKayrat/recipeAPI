package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/iKairat/RecipeAPI/internal/db/sqlc"
)

// Create a new recipe
func (s *Server) createRecipe(c *fiber.Ctx) error {
	req := new(requestBody)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	// decode steps into map
	ttime := make(map[string]int, 0)

	err = json.Unmarshal(req.Steps, &ttime)
	if err != nil {
		return err
	}

	// calculate total time of each step
	var totaltime int = 0
	for _, v := range ttime {
		totaltime += v
	}

	arg := db.CreateRecipeParams{
		Name:        req.Name,
		Description: req.Description,
		Ingredients: req.Ingredients,
		Steps:       req.Steps,
		TotalTime:   int16(totaltime),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// writtig recipe into DB
	recipe, err := s.store.CreateRecipe(c.Context(), arg)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(wrapMessage(err.Error()))
	}

	resp := response{}
	resp.Recipes = append(resp.Recipes, recipe)
	resp.Amount = len(resp.Recipes)
	resp.Status = http.StatusCreated

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// Get Recipe By ID
func (s *Server) getRecipe(c *fiber.Ctx) error {

	paramId := c.Params("id")

	id, err := strconv.ParseInt(paramId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(wrapMessage(msgInvalidParams))
	}

	// 	Recipes from DB
	recipe, err := s.store.GetRecipeByID(c.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(wrapMessage(msgNotFound))
		}
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	resp := response{}
	resp.Recipes = append(resp.Recipes, recipe)
	resp.Amount = len(resp.Recipes)
	resp.Status = http.StatusCreated

	return c.Status(fiber.StatusOK).JSON(resp)
}

// Get All Recipes
func (s *Server) getAllRecipes(c *fiber.Ctx) error {

	recipes, err := s.store.GetRecipes(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(wrapMessage(msgNotFound))
		}
		log.Println(err)
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	resp := response{}
	resp.Recipes = append(resp.Recipes, recipes...)
	resp.Amount = len(resp.Recipes)
	resp.Status = fiber.StatusOK

	return c.Status(fiber.StatusOK).JSON(resp)
}

// Get Recipe by ingredients
func (s *Server) getRecipesByIngredients(c *fiber.Ctx) error {

	query := c.Query("ingredients")

	ingredientList := strings.Split(query, ",")

	recipes, err := s.store.GetRecipes(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(wrapMessage(msgNotFound))
		}
		log.Println(err)
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	filteredRecipes := []db.Recipe{}

	// Bust of all recipes``
	for _, recipe := range recipes {
		// check if recipes inlude ingredients
		for _, v := range recipe.Ingredients {
			for _, i := range ingredientList {
				if strings.Contains(v, i) {
					filteredRecipes = append(filteredRecipes, recipe)
				}
			}
		}
	}

	resp := response{}
	resp.Recipes = append(resp.Recipes, filteredRecipes...)
	resp.Amount = len(resp.Recipes)
	resp.Status = fiber.StatusOK

	return c.Status(fiber.StatusOK).JSON(resp)
}

// Get Recipe By Total Time
func (s *Server) getRecipesByTotaltime(c *fiber.Ctx) error {

	// get recipes from DB
	recipes, err := s.store.GetRecipesByTime(c.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return c.Status(fiber.StatusNotFound).JSON(wrapMessage(msgNotFound))
		}
		log.Println(err)
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	resp := response{}
	resp.Recipes = append(resp.Recipes, recipes...)
	resp.Amount = len(resp.Recipes)
	resp.Status = fiber.StatusOK

	return c.Status(fiber.StatusOK).JSON(resp)

}

// Update Recipe by ID
func (s *Server) updateRecipe(c *fiber.Ctx) error {

	paramID := c.Params("id")
	req := new(requestBody)

	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(wrapMessage(msgInvalidParams))
	}

	err = c.BodyParser(req)
	if err != nil {
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	arg := db.UpdateRecipeParams{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Ingredients: req.Ingredients,
		Steps:       req.Steps,
		UpdatedAt:   time.Now(),
	}

	// Update in DB
	recipe, err := s.store.UpdateRecipe(c.Context(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return c.Status(http.StatusNotFound).JSON(wrapMessage(msgNotFound))
		}
		log.Println(err)
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	resp := response{}
	resp.Recipes = append(resp.Recipes, recipe)
	resp.Amount = len(resp.Recipes)
	resp.Status = fiber.StatusOK

	return c.Status(http.StatusAccepted).JSON(resp)
}

// Delete Recipe By ID
func (s *Server) deleteRecipe(c *fiber.Ctx) error {

	paramID := c.Params("id")

	id, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(wrapMessage(msgInvalidParams))
	}

	// Deleting in DB
	recipeID, err := s.store.DeleteRecipeByID(c.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(wrapMessage(msgNotFound))
		}
		return c.Status(503).JSON(wrapMessage(err.Error()))
	}

	message := fmt.Sprintf("deleted id - %d", recipeID)

	return c.Status(http.StatusAccepted).JSON(wrapMessage(message))
}
