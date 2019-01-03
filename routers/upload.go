package routers

import (
	"fmt"
	"golearning/models"
	u "golearning/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Uploadfile(c *gin.Context) {
	file, err := c.FormFile("file")
	userId, _ := strconv.Atoi(c.Query("userId"))
	if err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusBadRequest, err, "Upload file faild."))
		return
	}
	rename := fmt.Sprintf("%d.%s", time.Now().Unix(), strings.Split(file.Filename, ".")[1])
	dst := fmt.Sprintf("./upload/%s", rename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	var files models.Files
	files.UserId = userId
	files.Filename = file.Filename
	files.Realname = rename
	// if err := c.Bind(&files); err != nil {
	// 	c.AbortWithStatusJSON(u.Res(http.StatusBadRequest, err))
	// 	return
	// }
	if err := models.PostFile(files); err != nil {
		c.AbortWithStatusJSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, ""))
}
