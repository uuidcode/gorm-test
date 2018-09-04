package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
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

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func toJson(object interface{}) string {
	bytes, err := json.MarshalIndent(object, "", "    ")
	CheckErr(err)
	return string(bytes)
}

func main() {
	url := "root:rootroot@tcp(127.0.0.1:3306)/querydsl?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	CheckErr(err)

	db.LogMode(true)

	defer db.Close()

	name, err := uuid.NewV4()
	CheckErr(err)

	db.Create(&Book{
		Name:        name.String(),
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
