package middlewares

import (
	"github.com/Xavier577/hng-task-2/pkg/dtos"
	"github.com/gofiber/fiber/v2"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	// only bad exceptions are handled at the moment TODO: UPDATE TO HANDLE ALL ERROR CASES
	return c.Status(fiber.StatusBadRequest).JSON(dtos.ResponseBody{
		StatusCode: 400,
		Message:    err.Error(),
	})
}
