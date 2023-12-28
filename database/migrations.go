package database

import "golang-api/models"

func Migrate() {
    DB.AutoMigrate(&models.User{}, &models.Photo{})
}