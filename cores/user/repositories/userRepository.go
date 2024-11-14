package repositories

import "github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"

type UserRepository interface {
	GetById(id string) (*entities.User, error)
	Save(user *entities.User) error
	Delete(id string) error
}
