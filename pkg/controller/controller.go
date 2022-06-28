package controller

import (
	"practice/pkg/authentication"

	"github.com/gin-gonic/gin"
)



func RegisterRoutes(r *gin.Engine) {

	routes := r.Group("/api/v1/auth")
	routes.POST("/register", authentication.Register)

}


