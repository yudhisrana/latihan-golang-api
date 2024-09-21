package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yudhisrana/latihan-golang-api/controllers"
)

func RegisterUserRoute(app *fiber.App) {
	app.Get("/api/users", controllers.GetAllUsers)
	app.Get("/api/user/:id", controllers.GetUserById)
}
