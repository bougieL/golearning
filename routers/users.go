package routers

import (
	"golearning/models"
	u "golearning/utils"
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
	c.JSON(u.Res(http.StatusOK, users))
}

// GetUserByID func
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	// exist, _ := models.ExistUserByID(id)
	user, err := models.GetUserByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(u.Res(http.StatusOK, user))
}

// PostUser func
func PostUser(c *gin.Context) {
	var user models.Users
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusBadRequest, err))
		return
	}
	name := user.Username
	if exist, err := models.ExistUserByName(name); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusInternalServerError, err))
		return
	} else if exist {
		c.AbortWithStatusJSON(u.Res(http.StatusBadRequest, name+" Aleardy exists", name+" Aleardy exists"))
		return
	}
	if err := models.PostUser(user); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, "Add user success"))
}

// PutUserByID func
func PutUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var user models.Users
	if err := c.Bind(&user); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusBadRequest, err))
		return
	}
	name := user.Username
	if exist, err := models.ExistUserByNamePut(id, name); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusInternalServerError, err))
		return
	} else if exist {
		c.AbortWithStatusJSON(u.Res(http.StatusBadRequest, name+" Aleardy exists"))
		return
	}
	if err := models.PutUserByID(id, user); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, "Put user success"))
}

// DeleteUserByID func
func DeleteUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if err := models.DeleteUserByID(id); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, "Delete user success"))
}
