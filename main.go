package main

import (
	"golearning/models"

	"golearning/routers"
)

func main() {
	models.Setup()
	routers.InitRouters()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong pong",
	// 	})
	// })
	// r.Run()
}
