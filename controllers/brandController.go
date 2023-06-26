package controllers

import (
	"github.com/gin-gonic/gin"
)

func BrandCreate(c *gin.Context) {

	// post := models.Post{Title: "Hello World", Body: "Test"}

	// result := initializers.DB.Create(&post)

	// if result.Error != nil {
	// 	c.Status(400)
	// 	return
	// }

	c.JSON(200, gin.H{
		"message": "gets",
	})
}
