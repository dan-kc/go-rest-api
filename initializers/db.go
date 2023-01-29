package initializers

import (
	"fmt"
	"github.com/dan-kc/go-rest-api/packages/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDatabase() {

	var err error

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database", DB)
	}

}

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
}
