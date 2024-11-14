package server

import (
	"fmt"

	"github.com/Thnnathat/smart-todolist-backend-go/config"
	userHandlers "github.com/Thnnathat/smart-todolist-backend-go/cores/user/handlers"
	userRepositories "github.com/Thnnathat/smart-todolist-backend-go/cores/user/repositories"
	userUsecases "github.com/Thnnathat/smart-todolist-backend-go/cores/user/usecases"
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

	s.app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "OK",
		})
	})

	s.initializeUserHttpHandler()

	fmt.Printf("%s:%d", s.conf.Db.Host, s.conf.Server.Port)
	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Listen(serverUrl)
}

func (s *fiberServer) initializeUserHttpHandler() {

	api := s.app.Group("/api") // /api

	api = api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	// Initialize all layers
	userPostgresRepository := userRepositories.NewUserPostgresRepository(s.db)

	userUsecase := userUsecases.NewUserUsecaseImpl(userPostgresRepository)

	userHttpHandler := userHandlers.NewUserHttpHandler(userUsecase)

	//routers
	userRouter := api.Group("/users")
	userRouter.Post("", userHttpHandler.CreateUser)
	userRouter.Delete("/:id", userHttpHandler.DeleteUser)
}
