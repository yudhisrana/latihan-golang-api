package migrations

import (
	"log"

	"github.com/yudhisrana/latihan-golang-api/database"
	"github.com/yudhisrana/latihan-golang-api/models/entity"
)

func UserMigrate() {
	if err := database.DB.AutoMigrate(&entity.Users{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("sucsessfully migrated database users")
}