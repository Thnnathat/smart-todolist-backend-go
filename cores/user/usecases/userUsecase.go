package usecases

import "github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"

type UserUsecase interface {
	CreateUser(user *entities.User) error
	DeleteUser(id string) error
}
