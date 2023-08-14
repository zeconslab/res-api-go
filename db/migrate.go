package db

import (
	"github.com/zeconslab/res-api-go/models"
)

func Migrates() {
	DB.AutoMigrate(models.Task{})
	DB.AutoMigrate(models.User{})
}
