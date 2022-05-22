package config

import (
	"gin_sample_web_app/controller"

	"github.com/gin-gonic/gin"
)

func defineRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", controller.ShowAllTodo)
	router.POST("/new", controller.InsertTodo)
	router.GET("/detaile/:id", controller.ShowOneTodo)
	router.POST("/update/:id", controller.UpdateTodo)
	router.POST("/deleter/:id", controller.DeleteTodo)

	return router

}
