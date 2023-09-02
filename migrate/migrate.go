package main

import (
	"actors-backend/initializers"
	"actors-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Actor{})
}
