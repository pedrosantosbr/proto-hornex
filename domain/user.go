package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type User struct {
	ID          string    `json:"id" valid:"uuid" gorm:"type:uuid"`
	FirstName   string    `json:"firstName" valid:"notnull" gorm:"type:varchar(255)"`
	LastName    string    `json:"lastName" valid:"notnull" gorm:"type:varchar(255)"`
	DateOfBirth time.Time `json:"dateOfBirth" valid:"notnull" gorm:"type:datetime"` // ISO 8601
	Active      bool      `json:"active" valid:"notnull" gorm:"type:bool"`
	Email       string    `json:"email" valid:"notnull" gorm:"type:varchar(255);unique"`
	CreatedAt   time.Time `json:"createdAt" valid:"notnull" `
	UpdatedAt   time.Time `json:"updatedAt" valid:"notnull" `
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewUser() *User {
	return &User{}
}

func (u *User) Validate() error {

	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	return nil
}
