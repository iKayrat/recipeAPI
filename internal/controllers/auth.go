package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/iKairat/RecipeAPI/internal/db/sqlc"
	"github.com/iKairat/RecipeAPI/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (s *Server) register(c *fiber.Ctx) error {
	body := new(User)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(wrapMessage(err.Error()))
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	arg := db.CreateUserParams{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Username:  body.Username,
		Email:     body.Email,
		Password:  string(hashed),
	}

	created, err := s.store.CreateUser(c.Context(), arg)
	if err != nil {
		if strings.ContainsAny(err.Error(), "unique") {
			return c.Status(fiber.StatusInternalServerError).JSON(wrapMessage(msgUserAlreadyExists))
		}
		return c.Status(fiber.StatusInternalServerError).JSON(wrapMessage(err.Error()))

	}

	return c.JSON(created)
}

func (s *Server) login(c *fiber.Ctx) error {
	var body User

	if err := c.BodyParser(&body); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(wrapMessage(err.Error()))
	}

	user, err := s.store.GetUserByEmail(c.Context(), body.Email)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(wrapMessage(msgUserNotFound))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(wrapMessage(msgIncorrectPswd))
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(wrapMessage(msgSuccess))
}

func (s *Server) logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(wrapMessage(msgSuccess))

}
