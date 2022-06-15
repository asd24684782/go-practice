package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func GetHome(c *gin.Context) {

	c.String(http.StatusOK, time.Now().String())

}