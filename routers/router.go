package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"golearning/middleware"
)

// InitRouters func
func init() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("SESSION_CORE", store))
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
		apiUsers.DELETE("", middleware.Auth(), DeleteUserByID)
		apiUsers.POST("/login", Login)
		apiUsers.GET("/logout", Logout)
	}
	r.Run()
}
