package main

import (
	"go-practice/src/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", controller.GetHome)
	route.Run(":8080")
}