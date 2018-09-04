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

	defer db.Close()

	db.Create(&Book{
		Name:        uuid.NewV4().String(),
		UserId:      1,
		RegDatetime: time.Now(),
		ModDatetime: time.Now(),
	})

	rows, err := db.Raw("select * from book").Rows()
	CheckErr(err)

	defer rows.Close()

	var bookList []Book

	for rows.Next() {
		var book Book
		db.ScanRows(rows, &book)
		bookList = append(bookList, book)
	}

	fmt.Println(toJson(bookList))
}
