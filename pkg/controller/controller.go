package controller

import (
	"practice/pkg/api"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(r *gin.Engine) {

	routes := r.Group("/api/v1/auth")
	routes.POST("/register", api.Register)
	routes.GET("/find", api.FindUser)
}


