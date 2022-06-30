package api

import (
	"net/http"
	"practice/pkg/common/db"
	"practice/pkg/common/models"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

// @Summary     find All
// @Tags        user
// @version     1.0
// @produce     application/json
// @Success     200         {object}  models.User
// @Router      /api/v1/auth/find [get]
func FindUser(c *gin.Context) {

    var user models.User

	users := make([]models.User, 0)
    // 查詢資料列：
    println("查詢資料列：")
    rows, err := db.DB.Query("SELECT * FROM users")

    if err != nil {
        panic(err)
    }

    for rows.Next() {
			var uid *int
            err = rows.Scan(&uid, &user.Name, &user.Password, &user.Email)
			println(*uid)
			users = append(users, user)
			if err != nil {
				panic(err)
			}
    }
	//println(users)
    c.JSON(http.StatusCreated, &users)
}