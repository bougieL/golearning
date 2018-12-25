package routers

import (
	"golearning/models"

	"github.com/gin-gonic/gin"
)

// GetAllUsers func
func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		panic(err)
	}
	c.JSON(200, users)
}

// GetUserById func
// func GetUserById(c *gin.Context) {

// }
