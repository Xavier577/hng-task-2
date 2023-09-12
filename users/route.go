package users

import (
	"github.com/Xavier577/hng-task-2/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router) {
	r.Post("/", middlewares.RequestBodyValidatorMiddleware(new(CreateUserDTO)), createUser)

	r.Get("/:user_id", getUser)

	r.Patch("/:user_id", middlewares.RequestBodyValidatorMiddleware(new(UpdateUserDTO)), updateUser)

	r.Delete("/:user_id", deleteUser)

}
