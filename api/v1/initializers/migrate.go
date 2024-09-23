package initializers

import "github.com/kweku-xvi/go-users/api/v1/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
