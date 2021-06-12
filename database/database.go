package db

import (
	"fmt"

	modeluser "fibergo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Mydb() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost)/testgolang"), &gorm.Config{})

	if err != nil {
		panic("DB Cannot Connect")
	} else {
		fmt.Println("Connected")
	}

	DB = db

	DB.AutoMigrate(&modeluser.User{})

}
