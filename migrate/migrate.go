package main

import (
	"helloapp/internal/config"
	"helloapp/internal/domain/model"
	"helloapp/internal/initializers"
)

func init() {
	initializers.ConnectToDB()
	config.Load()

}

func main() {
	initializers.DB.AutoMigrate(&model.User{}, &model.Contact{})
}
