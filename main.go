package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// User routes
	app.Get("/users", getUsers)
	app.Post("/users", createUser)
	app.Get("/users/:id", getUser)

	// Start server
	app.Listen(":3000")
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func getUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	user.ID = len(users) + 1
	users = append(users, *user)
	return c.JSON(user)
}

func getUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	for _, user := range users {
		if user.ID == id {
			return c.JSON(user)
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "User not found")
}
