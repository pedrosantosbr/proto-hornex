package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pedrosantosbr/proto-hornex/domain"
)

type DB struct {
	Users []domain.User `json:"users"`
}

func (db *DB) Insert(user domain.User) error {
	db.Users = append(db.Users, user)
	return nil
}

func (db *DB) Update(user domain.User) error {
	for i, u := range db.Users {
		if u.ID == user.ID {
			db.Users[i] = user
			return nil
		}
	}

	return fmt.Errorf("user not found")
}

func (db *DB) GetAll() ([]domain.User, error) {
	return db.Users, nil
}

func NewDB() *DB {
	return &DB{
		Users: []domain.User{},
	}
}

func main() {
	app := fiber.New()

	db := NewDB()

	// Routers (Handlers)
	app.Post("/api/users", func(c *fiber.Ctx) error {
		var user domain.User

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		user.ID = len(db.Users) + 1

		db.Insert(user)

		fmt.Println(db.Users)

		return c.Status(http.StatusCreated).SendString("created")
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {

		users, _ := db.GetAll()

		return c.JSON(users)
	})

	app.Get("/api/users/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("GET /api/users/%s", c.Params("id")))
	})

	app.Delete("/api/users/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("GET /api/users/%s", c.Params("id")))
	})

	app.Put("/api/users/:id", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("PUT /api/users/%s", c.Params("id")))
	})

	log.Fatal(app.Listen(":9234"))
}
