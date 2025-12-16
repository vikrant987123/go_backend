package main

import(
	"github.com/gofiber/fiber/v2"
	"goBackend/internal/routes"
)

func main(){
	app:= fiber.New()

	// app.Get("/",func(c *fiber.Ctx) error{
	// 	return c.SendString("Go backend is running")
	// })

	routes.Register(app)

	app.Listen(":8080")
}