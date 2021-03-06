package controller

import (
	"fmt"
	"gin_sample_web_app/infrastructure/dbclient"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowAllTodo(ctx *gin.Context) {
	dbclient.DbInit()
	todos := dbclient.DbGetAll

	ctx.HTML(200, "index.html", gin.H{
		"todos": todos,
	})
}

func InsertTodo(ctx *gin.Context) {
	dbclient.DbInit()
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	dbclient.DbInsert(text, status)
	fmt.Println(111111)
	fmt.Println(ctx)
	ctx.Redirect(302, "/")
}

func ShowOneTodo(ctx *gin.Context) {
	dbclient.DbInit()
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	todo := dbclient.DbGetOne(id)
	ctx.HTML(200, "detail.html", gin.H{"todo": todo})
}

func UpdateTodo(ctx *gin.Context) {
	dbclient.DbInit()
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("Error")
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	dbclient.DbUpdate(id, text, status)
	ctx.Redirect(302, "/")
}

func DeleteTodo(ctx *gin.Context) {
	dbclient.DbInit()
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("Error")
	}
	dbclient.DbDelete(id)
	ctx.Redirect(302, "/")
}
