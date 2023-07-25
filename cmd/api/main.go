package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/render"
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

func (db *DB) Update(id int, params domain.User) error {

	for i, u := range db.Users {
		if u.ID == id {
			params.ID = id

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
	app := chi.NewRouter()

	db := NewDB()

	// Routers (Handlers)
	app.Post("/api/users", func(w http.ResponseWriter, r *http.Request) {
		var newUser = domain.User{Active: true}

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			return
		}

		defer r.Body.Close()

		for i := 0; i < len(db.Users); i++ {
			user := db.Users[i]

			if user.Email == newUser.Email {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, "This email is already in use.")
				return
			}
		}

		db.Insert(&newUser)

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, newUser)
	})

	app.Get("/api/users", func(w http.ResponseWriter, r *http.Request) {

		users, _ := db.GetAll()

		render.JSON(w, r, users)
	})

	/*
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
	*/

	srv := http.Server{Addr: ":9234", Handler: app}

	srv.ListenAndServe()

	log.Fatal()
}
