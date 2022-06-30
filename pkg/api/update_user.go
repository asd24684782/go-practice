package api

import (
	"net/http"
	"practice/pkg/common/db"
	"practice/pkg/common/models"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// @Summary     update user
// @Tags        user
// @version     1.0
// @produce     application/json
// @param       register    body      models.User true "User data"
// @Success     200         {object}  models.User
// @Router      /api/v1/user [put]
func UpdateUser(c *gin.Context) {

    var body models.User

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    // 更改資料列：
    println("\n更改資料列中...")
	stmt, err := db.DB.Prepare("UPDATE users SET password=$1 WHERE username=$2")

	if err != nil {
		panic(err)
 	}

    res, err := stmt.Exec(body.Password, body.Name)
	if err != nil {
		panic(err)
 	}

    affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
 	}
    println("已更改", affect, "個資料列。\n")
    c.JSON(http.StatusCreated, &body)
}