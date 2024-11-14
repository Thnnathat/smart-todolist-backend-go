package handlers

import "github.com/gofiber/fiber/v2"

type baseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func response(c *fiber.Ctx, responseCode int, message string, data interface{}) error {
	return c.Status(responseCode).JSON(&baseResponse{
		Code:    responseCode,
		Message: message,
		Data:    data,
	})
}
