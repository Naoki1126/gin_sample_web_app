package dbclient

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"gin_sample_web_app/model"
)

//gormはRailsでいうActiveRecordのような操作を可能にする

// DB初期化
func DbInit() {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db init function of main package")
	}

	db.AutoMigrate(&model.Todo{})
	// defer 上位の関数がreturnされるまで処理を遅延する
	defer db.Close()
}

//DBへRecordをInsertする
func DbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db insert function of main package")
	}

	db.Create(&model.Todo{Text: text, Status: status})
	defer db.Close()
}

//DBのRecordをUpdateする
func DbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db update function of main package")
	}

	var todo model.Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()

}

//DBのRecordを削除する
func DbDelete(id int) {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db delete function of main package")
	}

	var todo model.Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//Recordの全件取得
func DbGetAll() []model.Todo {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db get all function of main package")
	}

	var todos []model.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DBより1件のみ取得
func DbGetOne(id int) model.Todo {
	db := OpenDb()

	var todo model.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

//テスト用　後で使いたい
func OpenDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "gin_sample_web_app.sqlite3")
	if err != nil {
		panic("can not open database in db get all function of main package")
	}
	return db
}
