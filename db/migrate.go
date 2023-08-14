package db

import (
	"log"

	"github.com/zeconslab/res-api-go/models"
)

func Migrates() {
	log.Println("Migrating models...")
	DB.AutoMigrate(models.Task{})
	DB.AutoMigrate(models.User{})
	log.Println("Migrating models... OK")
}
