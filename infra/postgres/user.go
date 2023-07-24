package postgres

import "github.com/pedrosantosbr/proto-hornex/domain"

type UserModel struct {
	FirstName   string `gorm:"first_name" json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"` // ISO 8601
	Email       string `json:"email"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type PostgresUserRepositoryImpl struct {
}

func (r *PostgresUserRepositoryImpl) Insert(user domain.UserParams) (domain.User, error) {
	return domain.User{}, nil
}
