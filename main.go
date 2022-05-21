package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

//gormはRailsでいうActiveRecordのような操作を可能にする

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db init function of main package")
	}

	db.AutoMigrate(&Todo{})
	// defer 上位の関数がreturnされるまで処理を遅延する
	defer db.Close()
}

//DBへRecordをInsertする
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db insert function of main package")
	}

	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DBのRecordをUpdateする
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db update function of main package")
	}

	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()

}

//DBのRecordを削除する
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db delete function of main package")
	}

	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//Recordの全件取得
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db get all function of main package")
	}

	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DBより1件のみ取得
func dbGetOne(id int) Todo {
	db := openDb()

	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

//テスト用　後で使いたい
func openDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db get all function of main package")
	}
	return db
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	showText := "Hello Go/Gin"

	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"showText": showText,
			"todos":    todos,
		})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbInsert(text, status)
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		todo := dbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})

	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("Error")
		}
		todo := dbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	//Delete
	router.POST("/delete/:id")

	router.Run()
}
