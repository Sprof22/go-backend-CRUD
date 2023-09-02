package controllers

import (
	"actors-backend/initializers"
	"actors-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateActor(c *gin.Context) {
	//get Data of Req body
	var body struct {
		ActorName       string
		ActorRating     int
		ImagePath       string
		AlternativeName string
		ActorID         int
	}

	c.Bind(&body)

	//create actor
	actor := models.Actor{ActorName: body.ActorName, ActorRating: body.ActorRating, ImagePath: body.ImagePath, AlternativeName: body.AlternativeName, ActorID: body.ActorID}

	result := initializers.DB.Create(&actor)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	//return

	c.JSON(200, gin.H{
		"actor": actor,
	})
}

func GetActors(c *gin.Context) {
	var actors []models.Actor
	initializers.DB.Find(&actors)

	// if actors.Error != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": result.Error.Error(),
	// 	})
	// 	return
	// }

	//return

	c.JSON(200, gin.H{
		"actors": actors,
	})
}

func GetSingleActor(c *gin.Context) {

	// Get Id from the URL:
	id := c.Param("id")

	var actor models.Actor
	initializers.DB.First(&actor, id)

	c.JSON(200, gin.H{
		"actor": actor,
	})
}

func UpdateActor(c *gin.Context) {
	// Get Id from the URL:
	id := c.Param("id")

	var body struct {
		ActorName       string
		ActorRating     int
		ImagePath       string
		AlternativeName string
		ActorID         int
	}

	c.Bind(&body)

	var actor models.Actor
	initializers.DB.First(&actor, id)

	initializers.DB.Model(&actor).Updates(models.Actor{ActorName: body.ActorName, ActorRating: body.ActorRating, ImagePath: body.ImagePath, AlternativeName: body.AlternativeName, ActorID: body.ActorID})

	c.JSON(200, gin.H{
		"actor": actor,
	})
}

func DeleteActor(c *gin.Context) {
	// Get Id from the URL:
	id := c.Param("id")

	//get the particular person

	initializers.DB.Delete(&models.Actor{}, id)
	//return
	c.Status(200)
}
