package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"

	"goBackend/internal/models"
	"goBackend/internal/service"
)

var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "dob must be YYYY-MM-DD",
		})
	}

	age := service.CalcAge(dob)

	//eg.
	user:= models.User{
		ID: 1,
		Name: req.Name,
		DOB: dob,
		Age: age,
	}

	return c.Status(201).JSON(user)
}

func GetUser(c *fiber.Ctx) error {
	user := models.User{
		ID: 1,
		Name: "Vikrant",
		DOB: time.Date(2004,3,24,0,0,0,0, time.UTC),
	}

	user.Age = service.CalcAge(user.DOB)

	return c.JSON(user)
}