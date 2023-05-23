package main

import (
	"github.com/eputrap/GoCRUD/initializers"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadENVVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}