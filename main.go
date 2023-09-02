package main

import (
	"actors-backend/controllers"
	"actors-backend/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/actors", controllers.CreateActor)
	r.GET("/actors", controllers.GetActors)
	r.GET("/actors/:id", controllers.GetSingleActor)
	r.PUT("/actors/:id", controllers.UpdateActor)
	r.DELETE("/actors/:id", controllers.DeleteActor)
	r.Run()
}
