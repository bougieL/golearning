package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouters func
func InitRouters() {
	r := gin.Default()
	r.Use(cors.Default())
	apiUsers := r.Group("/api/user")
	apiUsers.GET("/all", GetAllUsers)
	r.Run()
}
