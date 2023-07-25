package services

import "time"

type UserCreateParams struct {
	ID       string `json:id`
	Email    string `json:id`
	Password string `json:id`
	Email    string `json:id`
	Email    string `json:id`

	DateOfBirth time.Time `json:id`
}
