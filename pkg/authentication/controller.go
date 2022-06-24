package authentication

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type handler struct {
    DB *sql.DB
}

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
    h := &handler{
        DB: db,
    }
	
	routes := r.Group("/auth")
	routes.POST("/register",h.Register)

}