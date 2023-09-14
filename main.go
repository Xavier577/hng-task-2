package main

import (
	"fmt"
	"github.com/Xavier577/hng-task-2/config/env"
	"github.com/Xavier577/hng-task-2/database/postgres"
	_ "github.com/Xavier577/hng-task-2/docs"
	"github.com/Xavier577/hng-task-2/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
)

// @title HNGx stage 2 task Api
// @version 1.0
// @description This is a sample swagger doc.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api

func setAppRoutes(a *fiber.App) {

	a.Get("/Yae4gahthiGeechie_health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	api := a.Group("/api")

	users.SetRoutes(api)
}

func init() {

	SSLMode := env.Get("PG_SSL_MODE")

	if SSLMode == "" {
		SSLMode = "disable"
	}

	err := postgres.Connect(&postgres.PgConnectCfg{
		DBName:   env.Get("PG_DATABASE"),
		Host:     env.Get("PG_HOST"),
		User:     env.Get("PG_USER"),
		Password: env.Get("PG_PASSWORD"),
		PORT:     env.Get("PG_PORT"),
		SSLMode:  SSLMode,
	})

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	app := fiber.New(fiber.Config{
		AppName: "fiberapp",
	})

	app.Use(requestid.New())

	app.Use(cors.New(cors.Config{}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${locals:requestid}] ${host} - ${latency} | ${status] | ${path} | ${method} | { headers -> {${reqHeaders} } | { body -> ${body} | { res -> ${resBody} }\n",
		TimeFormat: "02-Jan-2006 15:04",
	}))

	setAppRoutes(app)

	ADDRESS := fmt.Sprintf(":%s", env.Get("PORT"))

	err := app.Listen(ADDRESS)

	if err != nil {
		log.Fatal(err)
	}

}
