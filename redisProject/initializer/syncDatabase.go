package initializer

import (
	"github.com/biyoba1/redisProject/internal/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Person{}, &models.Order{}, &models.Product{})
}
