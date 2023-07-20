package application

import (
	"github.com/pedrosantosbr/proto-hornex/domain"
)

type UserService struct {
	userRepository domain.UserRepository
}

func (s UserService) RegisterNewUser(userParams domain.UserParams) (domain.User, error) {
	user, err := s.userRepository.Insert(userParams)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
