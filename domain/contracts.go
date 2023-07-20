package domain

type UserParams struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"` // ISO 8601
	Email       string `json:"email"`
	Password    string `json:"password"`
}

// func Validate(userParams UserParams) error {
// 	user := domain.User{
// 		Email:    p.Email,
// 		Password: p.Password,
// 	}

// 	if err := validation.Validate(&user); err != nil {
// 		return domain.WrapErrorf(err, domain.ErrorCodeInvalidArgument, "validation.Validate")
// 	}

// 	if !p.TermsAccepted {
// 		return validation.Errors{
// 			"termsAccepted": domain.NewErrorf(domain.ErrorCodeInvalidArgument, "Terms and conditions not accepted"),
// 		}
// 	}
// }

type UserRepository interface {
	Insert(user UserParams) (User, error)
}
