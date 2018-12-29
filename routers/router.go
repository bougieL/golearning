package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouters func
func InitRouters() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Go service is running!")
	})
	apiUsers := r.Group("/api/user")
	{
		apiUsers.GET("", GetUserById)
		apiUsers.GET("/all", GetAllUsers)
		apiUsers.POST("", PostUser)
		apiUsers.PUT("", PutUserById)
		apiUsers.DELETE("", DeleteUserById)
	}
	r.Run()
}
