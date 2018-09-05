package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/uuidcode/coreutil"
	"time"
)

type Project struct {
	ProjectId   int64  `gorm:"PRIMARY_KEY"`
	Name        string `gorm:"size:256"`
	RegDatetime time.Time
}

func (Project) TableName() string {
	return "project"
}

func main() {
	url := "root:rootroot@tcp(127.0.0.1:3306)/querydsl?charset=utf8&parseTime=true"
	db, err := gorm.Open("mysql", url)
	coreutil.CheckErr(err)

	db.LogMode(true)

	defer db.Close()

	project := Project{}
	db.DropTable(&project)
	//db.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ").CreateTable(&project)
	db.CreateTable(&project)

	db.Create(Project{
		Name: "😋",
	})

	db.Create(Project{
		Name: "😃",
	})

	db.Find(&project, Project{
		Name: "😋",
	})
}
