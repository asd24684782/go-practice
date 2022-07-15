package api

import (
	"net/http"
	"practice/pkg/common/db"
	"practice/pkg/common/models"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// @Summary     find All
// @Tags        user
// @version     1.0
// @produce     application/json
// @Success     200         {object}  models.User
// @Router      /api/v1/user [get]
func FindUser(c *gin.Context) {

    var name string

    // getting request's body
    if err := c.BindJSON(&name); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    // 查詢資料列：
    println("查詢資料列：")
	str1 := "SELECT * FROM users WHERE username = '"
	str2 := "'"
	sql := str1 + name + str2
    rows, err := db.DB.Query(sql)

    if err != nil {
        panic(err)
    }
	
	var user models.User
	users := make([]models.User, 0)

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