package users

import (
	"database/sql"
	"errors"
	"github.com/Xavier577/hng-task-2/database/postgres"
	"github.com/Xavier577/hng-task-2/pkg/dtos"
	"github.com/gofiber/fiber/v2"
	"log"
)

func getUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	var user User

	queryErr := postgres.Client().Get(&user, "SELECT id, name FROM users WHERE id = $1", userID)

	if queryErr != nil {
		log.Println(queryErr)

		ConnectionClosedError := errors.Is(queryErr, sql.ErrTxDone) || errors.Is(queryErr, sql.ErrNoRows)

		var errMsg string

		var statusCode int

		if ConnectionClosedError {
			statusCode = fiber.StatusBadRequest
			errMsg = "Something went wrong"
		} else {
			statusCode = fiber.StatusBadRequest
			errMsg = "Invalid user id"
		}

		return c.Status(statusCode).JSON(dtos.ResponseBody{
			StatusCode: statusCode,
			Message:    errMsg,
		})
	}

	return c.JSON(dtos.ResponseBody{StatusCode: fiber.StatusOK, Message: "success", Data: user})
}

func createUser(c *fiber.Ctx) error {
	reqPayload := new(CreateUserDTO)

	_ = c.BodyParser(reqPayload)

	var newlyCreatedUser User

	queryErr := postgres.Client().Get(&newlyCreatedUser, "INSERT INTO users (name) values ($1) RETURNING *", reqPayload.Name)

	if queryErr != nil {
		log.Fatal(queryErr)
	}

	return c.JSON(dtos.ResponseBody{StatusCode: fiber.StatusOK, Message: "success", Data: newlyCreatedUser})
}

func updateUser(c *fiber.Ctx) error {
	reqPayload := new(UpdateUserDTO)

	userID := c.Params("user_id")

	_ = c.BodyParser(reqPayload)

	var newlyCreatedUser User

	queryErr := postgres.Client().Get(&newlyCreatedUser, "UPDATE users SET name = $1 WHERE id = $2 RETURNING *", reqPayload.Name, userID)

	if queryErr != nil {
		log.Fatal(queryErr)
	}

	return c.JSON(dtos.ResponseBody{StatusCode: fiber.StatusOK, Message: "updated", Data: newlyCreatedUser})
}

func deleteUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	_, queryErr := postgres.Client().Exec("DELETE FROM users WHERE id = $1", userID)

	if queryErr != nil {
		log.Println(queryErr)

		ConnectionClosedError := errors.Is(queryErr, sql.ErrTxDone) || errors.Is(queryErr, sql.ErrNoRows)

		var errMsg string

		var statusCode int

		if ConnectionClosedError {
			statusCode = fiber.StatusBadRequest
			errMsg = "Something went wrong"
		} else {
			statusCode = fiber.StatusBadRequest
			errMsg = "Invalid user id"
		}

		return c.Status(statusCode).JSON(dtos.ResponseBody{
			StatusCode: statusCode,
			Message:    errMsg,
		})
	}

	return c.JSON(dtos.ResponseBody{StatusCode: fiber.StatusOK, Message: "success"})
}
