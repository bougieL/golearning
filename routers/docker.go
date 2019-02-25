package routers

import (
	"bytes"
	"fmt"
	u "golearning/utils"
	"net/http"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var endpoint = "unix:///var/run/docker.sock"
var client, err = docker.NewClient(endpoint)

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

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ping(c *gin.Context) {
	ws, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
	defer ws.Close()
	for {
		mt, message, _ := ws.ReadMessage()
		ws.WriteMessage(mt, message)
	}
}

func ExecCommand(c *gin.Context) {
	ws, _ := upGrader.Upgrade(c.Writer, c.Request, nil)
	defer ws.Close()
	containerID := "0248133e4476fcbc781ff061db7a2370e99403be1857bc8ebbb9cf901ecb29b4"

	execOption := docker.CreateExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		Container:    containerID,
		Cmd:          []string{"/bin/bash"},
	}
	execInstance, _ := client.CreateExec(execOption)
	fmt.Println(execInstance)
	for {
		mt, message, _ := ws.ReadMessage()
		var stdout, stdin bytes.Buffer
		stdin.Write(message)
		startExecOption := docker.StartExecOptions{
			Detach:       false,
			Tty:          false,
			RawTerminal:  true,
			InputStream:  &stdin,
			OutputStream: &stdout,
		}
		client.StartExec(execInstance.ID, startExecOption)
		fmt.Println(stdout.String())
		ws.WriteMessage(mt, []byte(stdout.String()))
	}
}

func ExecQuery(c *gin.Context) {
	var stdout, stdin, stderr bytes.Buffer
	command := c.Query("command")
	stdin.Write([]byte(command))
	containerID := "0248133e4476fcbc781ff061db7a2370e99403be1857bc8ebbb9cf901ecb29b4"
	execOption := docker.CreateExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		Container:    containerID,
		Cmd:          []string{"bash"},
	}
	execInstance, _ := client.CreateExec(execOption)
	startExecOption := docker.StartExecOptions{
		Detach:       false,
		Tty:          false,
		RawTerminal:  false,
		ErrorStream:  &stderr,
		InputStream:  &stdin,
		OutputStream: &stdout,
	}
	client.StartExec(execInstance.ID, startExecOption)
	c.JSON(200, gin.H{
		"stdin":  stdin.String(),
		"stdout": stdout.String(),
		"stderr": stderr.String(),
	})
}
