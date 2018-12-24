package routers

import (
	"github.com/gin-gonic/gin"
)

// InitRouters func
func InitRouters() {
	r := gin.Default()
	apiUsers := r.Group("/api/users")
	apiUsers.GET("/all", GetAllUsers)
	r.Run()
}
