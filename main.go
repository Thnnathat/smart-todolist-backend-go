package main

import (
	"github.com/Thnnathat/smart-todolist-backend-go/config"
	"github.com/Thnnathat/smart-todolist-backend-go/database"
	"github.com/Thnnathat/smart-todolist-backend-go/server"
)

func main() {
	conf := config.GetConfig()

	db := database.NewPostgresDatabase(conf)

	server.NewFiberServer(conf, db).Start()
}
