package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
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

func (db *DB) Find(id string) *domain.User {
	userId, _ := strconv.Atoi(id)

	for i := 0; i < len(db.Users); i++ {
		user := &db.Users[i]

		if userId == user.ID {
			return user
		}
	}

	return nil
}

func (db *DB) Deactivate(user *domain.User) {
	user.Active = false
}

/* func (db *DB) Update(id int, params domain.User) error {
	for i, u := range db.Users {
		if u.ID == id {
			params.ID = id
			db.Users[i] = params
			return nil
		}
	}

	return fmt.Errorf("user not found")
} */

func (db *DB) Update(id int, params domain.User) error {

	for i, u := range db.Users {
		if u.ID == id {
			params.ID = id
			// db.Users[i] = params

			mergeStruct[domain.User](&db.Users[i], params)
			return nil
		}
	}

	return fmt.Errorf("user not found")
}

func mergeStruct[T interface{}](m *T, toM T) {
	mVal := reflect.ValueOf(m).Elem()
	toMVal := reflect.ValueOf(toM)

	for i := 0; i < mVal.NumField(); i++ {
		mField := mVal.Field(i)
		toMField := toMVal.Field(i)

		if !toMField.IsZero() {
			mField.Set(toMField)
		}
	}
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
		userId := c.Params("id")

		userFound := db.Find(userId)

		if userFound == nil {
			return c.Status(http.StatusNotFound).SendString("User was not found.")
		}

		return c.JSON(userFound)
	})

	app.Delete("/api/users/:id", func(c *fiber.Ctx) error {
		userId := c.Params("id")

		userFound := db.Find(userId)

		if userFound == nil {
			return c.Status(http.StatusNotFound).SendString("User was not found.")
		}

		db.Deactivate(userFound)

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
