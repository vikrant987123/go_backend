package handler

import (
	
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	db "goBackend/db/sqlc"
	"goBackend/internal/models"
	"goBackend/internal/service"
)

var validate = validator.New()

func CreateUser(queries *db.Queries) fiber.Handler {
	return func (c *fiber.Ctx) error {
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

		user, err := queries.CreateUser(c.Context(), db.CreateUserParams{
			Name: req.Name,
			Dob: dob,
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error":"db error"})
		}

		return c.Status(201).JSON(user)
	}
}

func GetUser(queries *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _:= c.ParamsInt("id")

		user, err := queries.GetUserByID(c.Context(), int32(id))
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error":"user not found"})
		}

		age := service.CalcAge(user.Dob)

		return c.JSON(fiber.Map{
			"id": user.ID,
			"name": user.Name,
			"dob": user.Dob,
			"age": age,
		})
	}
}

func ListUsers(queries *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {

		page:= c.QueryInt("page",1)
		limit:= c.QueryInt("limit",10)

		if page<1 {
			page = 1
		}
		if limit<1 || limit>100 {
			limit = 10
		}

		offset := (page-1) * limit

		users, err := queries.ListUsers(
			c.Context(),
			db.ListUsersParams{
				Limit:  int32(limit),
				Offset: int32(offset),
			},
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":"db error",
			})
		}

		resp := []models.User{}
		for _, u := range users {
			resp = append(resp, models.User{
				ID:   int(u.ID),
				Name: u.Name,
				DOB:  u.Dob,
				Age:  service.CalcAge(u.Dob),
			})
		}

		return c.JSON(resp)
	}
}

func UpdateUser(queries *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")

		var req models.CreateUserRequest
		c.BodyParser(&req)

		dob, _ := time.Parse("2006-01-02", req.DOB)

		user, err := queries.UpdateUser(c.Context(), db.UpdateUserParams{
			ID:   int32(id),
			Name: req.Name,
			Dob:  dob,
		})
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "user not found"})
		}

		return c.JSON(user)
	}
}

func DeleteUser(queries *db.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")

		err:= queries.DeleteUser(c.Context(), int32(id))
		if err !=  nil {
			return c.Status(404).JSON(fiber.Map{"error":"user not found"})
		}

		return c.SendStatus(204)
	}
}