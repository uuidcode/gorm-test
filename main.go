package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/uuidcode/coreutil"
	"time"
)

type Book struct {
	BookId      int
	UserId      int
	Name        string
	RegDatetime time.Time
	ModDatetime time.Time
}

func (Book) TableName() string {
	return "book"
}

func toJson(object interface{}) string {
	bytes, err := json.MarshalIndent(object, "", "    ")
	coreutil.CheckErr(err)
	return string(bytes)
}

func main() {
	url := "root:rootroot@tcp(127.0.0.1:3306)/querydsl?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	coreutil.CheckErr(err)

	db.LogMode(true)

	defer db.Close()

	db.Create(&Book{
		Name:        uuid.NewV4().String(),
		UserId:      1,
		RegDatetime: time.Now(),
		ModDatetime: time.Now(),
	})

	var bookList []Book

	db.Find(&bookList, Book{
		UserId: 1,
	})

	fmt.Println(toJson(bookList))
}
