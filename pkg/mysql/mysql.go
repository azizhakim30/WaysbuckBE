package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit(){
	var err error
	// dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PORT"),os.Getenv("DB_NAME"))
	dsn := fmt.Sprintf("root@tcp(localhost:3306)/waysbuck?charset=utf8&parseTime=True&loc=Local")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println(("Connected to Database"))
}