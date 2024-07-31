package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Declare package-level DB variable

func ConnectToDB() {
	var err error
	dsn := "host=postgres user=root password=secret dbname=simple port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Assign to package-level DB variable
	if err != nil {
		panic("failed to connect database")
	}
	return
}
