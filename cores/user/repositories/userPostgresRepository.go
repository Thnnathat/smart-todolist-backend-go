package repositories

import (
	"strconv"

	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"
	"github.com/Thnnathat/smart-todolist-backend-go/database"
)

type userPostgresRepository struct {
	db database.Database
}

func NewUserPostgresRepository(db database.Database) UserRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) Save(user *entities.User) error {
	data := &entities.User{
		Username: user.Username,
		Password: user.Password,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *userPostgresRepository) Delete(id string) error {
	var userId, err = strconv.Atoi(id)

	if err != nil {
		return err
	}

	result := r.db.GetDb().Delete(&entities.User{}, uint(userId))

	if result.Error != nil {
		return err
	}

	return nil
}

func (r *userPostgresRepository) GetById(id string) (*entities.User, error) {
	var user entities.User
	idUint, err := gormIdConv(id)
	if err != nil {
		return nil, err
	}

	result := r.db.GetDb().Where("id = ?", idUint).Delete(&entities.User{})

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}
