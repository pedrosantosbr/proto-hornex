package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/go-chi/render"
	"github.com/pedrosantosbr/proto-hornex/domain"
	"github.com/pedrosantosbr/proto-hornex/framework/database"
)

func main() {
	db := database.NewDbTest()

	defer db.Close()

	app := chi.NewRouter()

	// Routers (Handlers)
	app.Post("/api/users", func(w http.ResponseWriter, r *http.Request) {
		var newUser = domain.User{Active: true}

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			return
		}

		render.Status(r, http.StatusCreated)
		render.JSON(w, r, newUser)
	})

	app.Get("/api/users", func(w http.ResponseWriter, r *http.Request) {
		users := []domain.User{}

		db.Find(&users)

		render.Status(r, http.StatusOK)
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

	log.Fatal(srv.ListenAndServe())

	defer srv.Close()
}
