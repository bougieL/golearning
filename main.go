package main

import (
	"golearning/models"

	"golearning/routers"
)

func main() {
	models.Setup()
	routers.InitRouters()
}
