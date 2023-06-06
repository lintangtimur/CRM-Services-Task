package main

import (
	"CRM-Services-Task/accounts"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/crmtask?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
func main() {
	r := gin.Default()
	var db, err = initDB()
	if err != nil {
		panic(err.Error())
	}

	arh := accounts.DefaultRequestHandler(db)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
