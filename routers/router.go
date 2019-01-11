package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouters func
func init() {
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
	// 	AllowCredentials: true,
	// }))
	r.Use(cors.Default())
	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Go service is running!")
	})
	apiFile := r.Group("/api/upload")
	{
		apiFile.POST("/file", Uploadfile)
	}
	apiUsers := r.Group("/api/user")
	{
		apiUsers.GET("", GetUserByID)
		apiUsers.GET("/all", GetAllUsers)
		apiUsers.POST("", PostUser)
		apiUsers.PUT("", PutUserByID)
		apiUsers.DELETE("", DeleteUserByID)
		apiUsers.POST("/login", Login)
	}
	r.Run()
}
