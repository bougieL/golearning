package middleware

import (
	"fmt"
	u "golearning/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Auth Middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.RequestURI)
		session := sessions.Default(c)
		userid := session.Get("userid")
		if userid != nil {
			c.Next()
		} else {
			c.JSON(u.Res(http.StatusUnauthorized, "User is not login."))
			c.Abort()
		}
	}
}
