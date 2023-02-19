package initializers

import (
	"fmt"
	"github.com/dan-kc/go-rest-api/packages/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	var err error

	dsn := "host=db user=postgres password=password dbname=go-rest-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database", DB)
	}

}

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
}
