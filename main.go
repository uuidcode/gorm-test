package main

import (
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

func (book Book) TableName() string {
	return "book"
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
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
}
