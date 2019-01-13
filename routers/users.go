package routers

import (
	"golearning/models"
	u "golearning/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

// Login func
func Login(c *gin.Context) {
	session := sessions.Default(c)
	var user models.Users
	if err := c.Bind(&user); err != nil {
		c.JSON(u.Res(http.StatusBadRequest, user))
		return
	}
	user.Password = u.Md5(user.Password)
	if userData, err := models.ValidateUserLogin(user); err != nil {
		c.JSON(u.Res(http.StatusBadRequest, err, "Invalidate username or password"))
	} else {
		session.Set("userid", userData.ID)
		session.Set("username", userData.Username)
		session.Save()
		c.JSON(u.Res(http.StatusOK, userData))
	}
}

// Logout func
func Logout(c *gin.Context) {
	location := c.DefaultQuery("location", strings.Split(c.Request.RequestURI, ":")[0])
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusTemporaryRedirect, location)
}

// GetAllUsers func
func GetAllUsers(c *gin.Context) {
	if users, err := models.GetAllUsers(); err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
	} else {
		c.JSON(u.Res(http.StatusOK, users))
	}
}

// GetUserByID func
func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	if user, err := models.GetUserByID(id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(u.Res(http.StatusOK, user))
	}
}

// PostUser func
func PostUser(c *gin.Context) {
	var user models.Users
	if err := c.Bind(&user); err != nil {
		c.JSON(u.Res(http.StatusBadRequest, err))
		return
	}
	name := user.Username
	if exist, err := models.ExistUserByName(name); err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	} else if exist {
		c.JSON(u.Res(http.StatusBadRequest, name+" Aleardy exists", name+" Aleardy exists"))
		return
	}
	user.Password = u.Md5(user.Password)
	if err := models.PostUser(user); err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, "Add user success"))
}

// PutUserByID func
func PutUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var user models.Users
	if err := c.Bind(&user); err != nil {
		c.JSON(u.Res(http.StatusBadRequest, err))
		return
	}
	name := user.Username
	if exist, err := models.ExistUserByNamePut(id, name); err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	} else if exist {
		c.JSON(u.Res(http.StatusBadRequest, name+" Aleardy exists"))
		return
	}
	if err := models.PutUserByID(id, user); err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, "Put user success"))
}

// DeleteUserByID func
func DeleteUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if err := models.DeleteUserByID(id); err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, "Delete user success"))
}
