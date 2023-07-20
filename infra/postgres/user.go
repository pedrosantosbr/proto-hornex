package postgres

type UserModel struct {
	FirstName   string `gorm:"first_name" json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"` // ISO 8601
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
