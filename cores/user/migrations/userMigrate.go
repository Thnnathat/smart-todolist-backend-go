package migrations

import (
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"
	"github.com/Thnnathat/smart-todolist-backend-go/database"
)

func UserMigrate(db database.Database) {
	db.GetDb().AutoMigrate(&entities.User{})
}
