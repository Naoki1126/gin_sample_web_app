package dbclient

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name         string `gorm:"size:255"`
	Age          int
	Sex          string `gorm:"size:255"`
	MessageToken string `gorm:"size:255"`
}

type MysqlSumResult struct {
	AgeSum int
}

func (u User) String() string {
	return fmt.Sprint("%s(%d)", u.Name, u.Age)
}

const (
	Dialect = "mysql"

	DBUser = "root"

	DBProtocol = "tcp(127.0.0.1:3306)"

	DBName = "gin_sample_web_app"

	Option = "parseTime=true"
)

func ConnectGorm() *gorm.DB {
	connectTemplate := "%s@%s/%s?%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBProtocol, DBName, Option)
	db, err := gorm.Open(Dialect, connect)
	// db.AutoMigrate(&User{})
	if err != nil {
		log.Println("Connected Error")
		log.Println(err.Error())
	}
	return db
}

func Insert(user User, db *gorm.DB) {
	db.NewRecord(user)
	db.Create(&user)
}

func FindAll(db *gorm.DB) []User {
	var allUsers []User
	db.Find(&allUsers)
	return allUsers
}

func FindLast(db *gorm.DB) (User, error) {
	var user User
	err := db.Last(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("record not found")
		return user, nil
	}

	// var user User
	fmt.Println(11)
	return user, nil
}
