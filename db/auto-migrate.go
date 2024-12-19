package db

import (
	"github.com/channyeintun/golang-jwt-authentication/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
