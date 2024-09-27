package initializers

import "github.com/orest-kostiuk/go-jwt/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
		return
	}
}
