package routers

import (
	"fmt"
	u "golearning/utils"
	"net/http"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

var endpoint = "unix:///var/run/docker.sock"
var client, err = docker.NewClient(endpoint)
var exec = client.CreateExec

func GetImages(c *gin.Context) {
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	imgs, err := client.ListImages(docker.ListImagesOptions{All: true})
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, imgs))
}

func SearchImage(c *gin.Context) {
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	keyword := c.Query("keyword")
	fmt.Println(keyword)
	imgs, err := client.SearchImages(keyword)
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, imgs))
}

func GetContainers(c *gin.Context) {
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, containers))
}

func ExecCommand(c *gin.Context) {
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	containerID := c.Query("containerid")
	command := c.Query("command")
	execOption := docker.CreateExecOptions{
		Container: containerID,
		Cmd:       []string{command},
	}
	stdout, err := client.CreateExec(execOption)
	if err != nil {
		c.JSON(u.Res(http.StatusInternalServerError, err))
		return
	}
	c.JSON(u.Res(http.StatusOK, stdout))
}
