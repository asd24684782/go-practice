package main

import (
	"practice/pkg/common/db"
	"practice/pkg/controller"

	_ "practice/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
    viper.SetConfigFile("../pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    //init db
    db.Init(dbUrl)
    controller.RegisterRoutes(r)

    url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
    
    r.Run(port)
}

/*
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

       // 初始化資料庫：
       db, err := sql.Open("postgres", "user=postgres password=5598 dbname=postgres sslmode=disable")
       defer db.Close()
       checkErr(err)

       // 新增資料列：
       stmt, err := db.Prepare("INSERT INTO auth(name,password) VALUES($1,$2);")
       checkErr(err)
       res, err := stmt.Exec("apple", "iphone哀鳳")
       checkErr(err)
       println("已新增資料列。\n")
       // 查詢資料列：
       println("查詢資料列：")
       rows, err := db.Query("SELECT * FROM auth")
       checkErr(err)
       var id, use, pw string
       for rows.Next() {
              err = rows.Scan(&id, &use, &pw)
              checkErr(err)
              fmt.Println("\t", use, pw)
       }

       // 更改資料列：
       println("\n更改資料列中...")
       stmt, err = db.Prepare("UPDATE auth SET password=$1 WHERE name=$2")
       checkErr(err)
       res, err = stmt.Exec("ipad2", "apple")
       checkErr(err)
       affect, err := res.RowsAffected()
       checkErr(err)
       println("已更改", affect, "個資料列。\n")

       // 查詢資料列：
       println("查詢資料列：")
       rows, err = db.Query("SELECT * FROM auth")
       checkErr(err)
       for rows.Next() {
              err = rows.Scan(&id, &use, &pw)
              checkErr(err)
              fmt.Println("\t", use, pw)
       }


       // 删除資料列：
       stmt, err = db.Prepare("DELETE FROM auth WHERE name=$1")
       checkErr(err)
       res, err = stmt.Exec("apple")
       checkErr(err)
       affect, err = res.RowsAffected()
       checkErr(err)
       println("\n已刪除", affect, "個資料列。\n")

       // 查詢資料列：
       println("查詢資料列：")
       rows, err = db.Query("SELECT * FROM auth")
       checkErr(err)
       if(!rows.Next()){
              print("資料列是空的 !!")
       }
}




func checkErr(err error) {
       if err != nil {
              panic(err)
       }
}*/
