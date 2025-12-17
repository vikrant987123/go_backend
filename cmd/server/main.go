package main

import(
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	db "goBackend/db/sqlc"
	"goBackend/internal/routes"
)

func main(){
	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:superuser@localhost:5432/go_backend?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	app:= fiber.New()

	// app.Get("/",func(c *fiber.Ctx) error{
	// 	return c.SendString("Go backend is running")
	// })

	routes.Register(app, queries)

	log.Fatal(app.Listen(":8080"))
}