package routers

import (
	"golearning/models"
	"net/http"
	"strconv"

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
func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	// exist, _ := models.ExistUserById(id)
	user, err := models.GetUserById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var user models.Users
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}
	if err := models.PostUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Add user success",
	})
}

func PutUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var user models.Users
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}
	if err := models.PutUserById(id, user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Put user success",
	})
}

func DeleteUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if err := models.DeleteUserById(id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete user success",
	})
}
