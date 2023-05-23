package main

import (
	"github.com/eputrap/GoCRUD/initializers"
	"github.com/eputrap/GoCRUD/models"
)

// create postgres table

func init() {
	initializers.LoadENVVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}