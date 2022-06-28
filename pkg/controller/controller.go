package controller

import (
	"database/sql"

	"practice/pkg/authentication"
	"github.com/gin-gonic/gin"
)

type Handler struct {
    DB *sql.DB
	
}


func RegisterRoutes(r *gin.Engine, db *sql.DB) {
    h := &Handler{
        DB: db,
    }
	
	routes := r.Group("/api/v1/auth")
	routes.POST("/register", h.Register)

}


