package server

import (
	"fmt"

	"github.com/Thnnathat/smart-todolist-backend-go/config"
	"github.com/Thnnathat/smart-todolist-backend-go/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type fiberServer struct {
	app  *fiber.App
	db   database.Database
	conf *config.Config
}

func NewFiberServer(conf *config.Config, db database.Database) Server {
	fiberApp := fiber.New()

	return &fiberServer{
		app:  fiberApp,
		db:   db,
		conf: conf,
	}
}

func (s *fiberServer) Start() {
	s.app.Use(cors.New())
	s.app.Use(logger.New())

	s.app.Get("v1/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	})

	fmt.Printf("%s:%d", s.conf.Db.Host, s.conf.Server.Port)
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Listen(serverUrl)
}
