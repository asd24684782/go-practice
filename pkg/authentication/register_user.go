package authentication

import (
	"fmt"
	"net/http"

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
// @param       register    body      models.Register true "register data"
// @Success     200         {object}  models.Register
// @Router      /api/v1/auth/register [post]
func (h handler) Register(c *gin.Context) {

    body := RegisterRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var register models.Register
	
	register.Name = body.Name
	register.Password = body.Password
    register.Email = body.Email

    fmt.Println(register)
    // 新增資料列：
    stmt, err := h.DB.Prepare("INSERT INTO users(username, password, email) VALUES($1, $2, $3);")
    if err != nil {
        panic(err)
    }

    res, err := stmt.Exec(register.Name, register.Password, register.Email)

    if err != nil {
        panic(err)
    }
    println("已新增資料列。\n",res)

    c.JSON(http.StatusCreated, &register)
}