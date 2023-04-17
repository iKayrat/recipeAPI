package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iKairat/RecipeAPI/internal/controllers"
	db "github.com/iKairat/RecipeAPI/internal/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {

	dbsource := os.Getenv("DBSOURCE")

	conn, err := sql.Open("postgres", dbsource)
	if err != nil {
		log.Println(err)
	}

	store := db.NewStore(conn)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	server := controllers.New(store, app)

	err = server.App.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
