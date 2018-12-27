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
	apiUsers.GET("", GetUserById)
	apiUsers.GET("/all", GetAllUsers)
	apiUsers.POST("", PostUser)
	apiUsers.PUT("", PutUserById)
	apiUsers.DELETE("", DeleteUserById)
	r.Run()
}
