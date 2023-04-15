package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iKairat/RecipeAPI/internal/controllers"
	db "github.com/iKairat/RecipeAPI/internal/db/sqlc"
	"github.com/iKairat/RecipeAPI/internal/utils"
	_ "github.com/lib/pq"
)

func main() {
	utils.InitPopulation()

	dbsource := os.Getenv("DBSOURCE")
	if dbsource == "" {
		dbsource = "postgresql://root:kaak@0.0.0.0:5432/recipe?sslmode=disable"
	}

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
