package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yudhisrana/latihan-golang-api/database"
	"github.com/yudhisrana/latihan-golang-api/database/migrations"
	"github.com/yudhisrana/latihan-golang-api/routes"
)

func main() {
	log.Println("server started on port 3000")
	database.ConnectDB()
	migrations.UserMigrate()

	app := fiber.New()
	routes.RegisterUserRoute(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}