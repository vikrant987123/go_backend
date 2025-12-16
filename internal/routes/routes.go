package routes

import(
	"github.com/gofiber/fiber/v2"
	"goBackend/internal/handler"
)

func Register(app *fiber.App){
	app.Get("users/:id", handler.GetUser)
	app.Post("/users",handler.CreateUser)
}