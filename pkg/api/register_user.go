package api

import (
	"net/http"
	"practice/pkg/common/db"
	"practice/pkg/common/models"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type RegisterRequestBody struct {
    Name       	string `json:"name"`
    Password    string `json:"password"`
    Email		string `json:"email"`
}


// @Summary     register
// @Tags        user
// @version     1.0
// @produce     application/json
// @param       register    body      models.User true "User data"
// @Success     200         {object}  models.User
// @Router      /api/v1/user [post]
func Register(c *gin.Context) {

    body := RegisterRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var user models.User
	
	user.Name = body.Name
	user.Password = body.Password
    user.Email = body.Email

    // 新增資料列：
    stmt, err := db.DB.Prepare("INSERT INTO users(username, password, email) VALUES($1, $2, $3);")
    if err != nil {
        panic(err)
    }

    res, err := stmt.Exec(user.Name, user.Password, user.Email)

    if err != nil {
        panic(err)
    }
    println("已新增資料列。\n",res)

    c.JSON(http.StatusCreated, &user)
}