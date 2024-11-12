package main

import (
	"github.com/Thnnathat/smart-todolist-backend-go/config"
	userMigration "github.com/Thnnathat/smart-todolist-backend-go/cores/user/migrations"
	"github.com/Thnnathat/smart-todolist-backend-go/database"
	"github.com/Thnnathat/smart-todolist-backend-go/server"
)

func main() {
	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)

	userMigration.UserMigrate(db)

	server.NewFiberServer(conf, db).Start()
}
