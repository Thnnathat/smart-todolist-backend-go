package handlers

import (
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/entities"
	"github.com/Thnnathat/smart-todolist-backend-go/cores/user/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type userHttpHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHttpHandler(userUsecase usecases.UserUsecase) UserHandler {
	return &userHttpHandler{userUsecase: userUsecase}
}

func (h *userHttpHandler) CreateUser(c *fiber.Ctx) error {
	reqBody := new(entities.User)

	if err := c.BodyParser(reqBody); err != nil {
		log.Errorf("Error parse request body: %v", err)
		return response(c, fiber.StatusBadRequest, "Bad request", reqBody)
	}

	if err := h.userUsecase.CreateUser(reqBody); err != nil {
		return response(c, fiber.StatusInternalServerError, "Create user failed", reqBody)
	}

	return response(c, fiber.StatusCreated, "Success", reqBody)
}
