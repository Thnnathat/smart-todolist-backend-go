package usecases

import (
	"github.com/Thnnathat/smart-todolist-backend-go/cores/errors"
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/repositories"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	createUser.Password = string(hashedPassword)

	if err := u.userRepository.Save(createUser); err != nil {
		return err
	}

	return nil
}

func (u *userUsecaseImpl) DeleteUser(id string) error {
	user, err := u.userRepository.GetById(id)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.ErrNotfound
	}

	if err := u.userRepository.Delete(id); err != nil {
		return err
	}

	return nil
}
