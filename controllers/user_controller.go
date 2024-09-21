package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yudhisrana/latihan-golang-api/database"
	"github.com/yudhisrana/latihan-golang-api/models/entity"
	"gorm.io/gorm"
)

// fungsi untuk mendapatkan semua data user
func GetAllUsers(c *fiber.Ctx) error {
	users := []entity.Users{} // deklarasi variable user

	if err := database.DB.Find(&users).Error; err != nil {
		// jika data tidak ditemukan
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "fail",
				"message": "No users found",
			})
		}

		// jika terjadi kesalahan lain 
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	// jika data ditemukan
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": users,
	})
}

// fungsi untuk mendapatkan data user berdasarkan id
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id") // ambil id dari parameter url
	user := []entity.Users{} // deklarasi variable user

	if err := database.DB.First(&user, id).Error; err != nil {
		// jika data tidak ditemukan
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": "fail",
				"message": "No user found with ID " + id,
			})
		}

		// jika terjadi kesalahan lain
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": err.Error(),
		})
	}

	// jika data ditemukan
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": user,
	})
}