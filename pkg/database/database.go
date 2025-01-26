package databse

import (
	"log"
	"todo/config"
	"todo/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := config.Get("DB_URL", "")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Faild to connect to db!")
		return
	}
	DB.AutoMigrate(&models.User{}, &models.Todo{})

}
