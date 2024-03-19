package main

import (
	"log"

	// "testing"
	api "project-fe-be/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Put("/api/put", api.Put)
	app.Patch("/api/patch", api.Patch)
	app.Get("/api/get", api.Get_Data)
	app.Post("/api/post", api.Post_Data)
	app.Delete("/api/delete", api.Delete)
	app.Listen(":8000")
	log.Print("Api is running")
}
