package repositories

import (
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"
	"github.com/Thnnathat/smart-todolist-backend-go/database"
	"github.com/gofiber/fiber/v2/log"
)

type userPostgresRepository struct {
	db database.Database
}

func NewUserPostgresRepository(db database.Database) UserRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) SaveUser(user *entities.User) error {
	data := &entities.User{
		Username: user.Username,
		Password: user.Password,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("Save User Data: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}
