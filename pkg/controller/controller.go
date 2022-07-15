package controller

import (
	"practice/pkg/api"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(r *gin.Engine) {

	routes := r.Group("/api/v1/user")
	routes.POST("/", api.Register)
	routes.GET("/", api.FindUsers)
	routes.PUT("/", api.UpdateUser)
	routes.POST("/login", api.Login)
}


