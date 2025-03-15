package db

import (
	"log"
	"os"

	"github.com/Frhnmj2004/FarmQuest-server.git/config"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Postgres.Url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	db.AutoMigrate(&models.User{}, &models.Crop{})
	return db
}
