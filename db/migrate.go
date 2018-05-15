package db

import (
	// "fmt"
	// "io/ioutil"
	"chat/model"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.LogMode(true)
	db.AutoMigrate(&model.User{})
	// 本来想用目录结构来做的
	// files, _ := ioutil.ReadDir("./model/")  
	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }
}