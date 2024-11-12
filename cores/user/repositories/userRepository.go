package repositories

import "github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"

type UserRepository interface {
	SaveUser(user *entities.User) error
}
