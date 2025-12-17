package main

import(
	"database/sql"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	db "goBackend/db/sqlc"
	"goBackend/internal/routes"
	"goBackend/internal/logger"
	"goBackend/internal/middleware"
)

func main(){
	if err := godotenv.Load();
	err != nil {
		log.Println("no .env")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	app:= fiber.New()

	// app.Get("/",func(c *fiber.Ctx) error{
	// 	return c.SendString("Go backend is running")
	// })

	logg := logger.New()

	app.Use(middleware.ResquestID())
	app.Use(middleware.Logger(logg))

	routes.Register(app, queries)

	log.Fatal(app.Listen(":8080"))
}