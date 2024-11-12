package usecases

import (
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/repositories"
)

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewUserUsecaseImpl(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}

func (u *userUsecaseImpl) CreateUser(user *entities.User) error {
	createUser := &entities.User{
		Username: user.Username,
		Password: user.Password,
	}

	if err := u.userRepository.SaveUser(createUser); err != nil {
		return err
	}

	return nil
}