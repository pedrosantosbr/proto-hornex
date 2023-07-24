package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pedrosantosbr/proto-hornex/domain"
)

type DB struct {
	Users []domain.User `json:"users"`
}

func (db *DB) Insert(user *domain.User) error {
	user.ID = len(db.Users) + 1
	db.Users = append(db.Users, *user)
	return nil
}

func (db *DB) Find(id string) domain.User {
	userId, _ := strconv.Atoi(id)
	return db.Users[userId]
}

func (db *DB) Update(id int, params domain.User) error {
	for i, u := range db.Users {
		if u.ID == id {
			params.ID = id
			db.Users[i] = params
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
		var newUser = domain.User{Active: true}

		if err := c.BodyParser(&newUser); err != nil {
			return err
		}

		for i := 0; i < len(db.Users); i++ {
			user := db.Users[i]

			if user.Email == newUser.Email {
				return c.Status(http.StatusBadRequest).SendString("This email is already in use.")
			}
		}

		db.Insert(&newUser)

		return c.Status(http.StatusCreated).JSON(newUser)
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {

		users, _ := db.GetAll()

		return c.JSON(users)
	})

	app.Get("/api/users/:id", func(c *fiber.Ctx) error {
		userId, _ := strconv.Atoi(c.Params("id"))

		var userFound domain.User

		for i := 0; i < len(db.Users); i++ {
			user := db.Users[i]

			if userId == user.ID {
				userFound = user
			}
		}

		if userFound.ID == 0 {
			return c.Status(http.StatusNotFound).SendString("User was not found.")
		}

		return c.JSON(userFound)
	})

	app.Delete("/api/users/:id", func(c *fiber.Ctx) error {
		userId, _ := strconv.Atoi(c.Params("id"))

		var userFound *domain.User

		for i := 0; i < len(db.Users); i++ {
			userFound = &db.Users[i]

			if userId == userFound.ID {
				userFound.Active = false
			}
		}

		if userFound.ID == 0 {
			return c.Status(http.StatusNotFound).SendString("User was not found.")
		}

		return c.SendStatus(http.StatusNoContent)
	})

	app.Put("/api/users/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var req domain.User

		if err := c.BodyParser(&req); err != nil {
			return err
		}

		if err := db.Update(id, req); err != nil {
			return c.Status(http.StatusNotFound).SendString("User was not found.")
		}

		return c.SendStatus(http.StatusNoContent)
	})

	log.Fatal(app.Listen(":9234"))
}
