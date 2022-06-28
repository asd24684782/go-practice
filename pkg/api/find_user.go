package api

import (
	"net/http"
	"practice/pkg/common/db"
	"practice/pkg/common/models"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// @Summary     register
// @Tags        user
// @version     1.0
// @produce     application/json
// @param       register    body      models.Register true "register data"
// @Success     200         {object}  models.Register
// @Router      /api/v1/auth/register [post]
func FindUser(c *gin.Context) {

    var user models.User
	users := make(map[string]models.User)

    // 查詢資料列：
    println("查詢資料列：")
    rows, err := db.DB.Query("SELECT * FROM users")

    if err != nil {
        panic(err)
    }

   

    for rows.Next() {
            err = rows.Scan(&user.Name, &user.Password, &user.Email)
			println(&user)
			users[user.Name] = user
			if err != nil {
				panic(err)
			}
    }
	//println(&users)
    c.JSON(http.StatusCreated, &user)
}