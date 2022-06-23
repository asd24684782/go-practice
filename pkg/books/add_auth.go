package books

import (
	"net/http"

	"go-practice/pkg/common/models"

	"github.com/gin-gonic/gin"
)

type AddAuthRequestBody struct {
    Name       	string `json:"name"`
    Password    string `json:"password"`
}

func (h handler) Authenticator(c *gin.Context) {
    body := AddAuthRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var auth models.Auth
	
	auth.Name = body.Name
	auth.Password = body.Password


    if result := h.DB.Create(&auth); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &auth)
}