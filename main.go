package main

import (
	// "golearning/models"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// models.Setup()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db, err := gorm.Open("mysql", "root:root@tcp(10.10.0.2:3306)/golearning?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalf("models.Setup err: %v", err)
		} else {
			log.Println("models.Setup success", db)
		}
		c.JSON(200, gin.H{
			"message": "pong pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
