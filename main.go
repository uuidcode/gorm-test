package main

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorm-test/logger"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/uuidcode/coreutil"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
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

func init() {
	log.SetFormatter(&prefixed.TextFormatter{})
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Debug("main")

	url := "root:rootroot@tcp(127.0.0.1:3306)/querydsl?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	coreutil.CheckErr(err)

	db.LogMode(true)
	db.SetLogger(logger.New())

	defer db.Close()

	log.Warn("connect")

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

	log.Debug(toJson(bookList))
}
