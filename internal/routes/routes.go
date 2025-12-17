package routes

import (
	"github.com/gofiber/fiber/v2"

	db "goBackend/db/sqlc"
	"goBackend/internal/handler"
)

func Register(app *fiber.App, queries *db.Queries) {
	app.Post("/users", handler.CreateUser(queries))
	app.Get("/users", handler.ListUsers(queries))
	app.Get("/users/:id", handler.GetUser(queries))
	app.Put("/users/:id", handler.UpdateUser(queries))
	app.Delete("/users/:id", handler.DeleteUser(queries))
}
