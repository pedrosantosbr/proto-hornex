package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CreateUserRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateOfBirth string `json:"dateOfBirth"`
	Password    string `json:"password"`
}

type CreateUserResponse struct {
	ID uint8 `json:"email"`
}

func main() {

	router := chi.NewRouter()

	router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("POST /users", r.Body)
		// var req CreateUserRequest
		// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// 	renderErrorResponse(w, r, "invalid request",
		// 		internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder"))

		// 	return
		// }

		render.Status(r, http.StatusOK)
		render.JSON(w, r, CreateUserResponse{ID: 1})
	})

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Println(err)
	}
}
