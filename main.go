package main

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"os"

	"github.com/go-eyas/mock-server/handler"
)

func initRouter(app *gin.Engine) {

	// admin UI
	app.Static("/client", "./client")

	// admin api
	admin := app.Group("/admin")
	admin.GET("/project", handler.GetProjects)
	admin.POST("/project", handler.CreateORUpdateProject)
	admin.DELETE("project", handler.DeleteProject)
}

func main() {
	db, err := handler.Connect("root:pwd+sql@tcp(localhost:3306)/mock-server?loc=Local&parseTime=True&charset=utf8mb4")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	db.AutoMigrate(&handler.APIS{})
	// db.Create(&handler.APIS{
	// 	Method: "get",
	// 	Path:   "/say",
	// 	Value:  "{}",
	// })
	app := gin.Default()
	app.Use(handler.Cors)
	app.Use(handler.APIHandler)
	initRouter(app)
	app.Run(":8001")
}
