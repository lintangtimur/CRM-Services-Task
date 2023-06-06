package main

import (
	"CRM-Services-Task/accounts"
	"CRM-Services-Task/customers"
	"CRM-Services-Task/middleware"
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
	crh := customers.DefaultRequestHandler(db)

	r.POST("/actors/login", arh.ActorLogin)

	protected := r.Group("")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.POST("/customers/create", crh.CreateCustomer)
	protected.POST("/actors/create", arh.CreateAdmin)
	protected.GET("/actors/approve", middleware.SuperadminMiddleware(), arh.GetApprove)
	protected.PUT("/actors/approve", middleware.SuperadminMiddleware(), arh.Approve)
	protected.PUT("/actors/reject", middleware.SuperadminMiddleware(), arh.Reject)
	protected.PUT("/actors/activate", middleware.SuperadminMiddleware(), arh.Activate)
	protected.PUT("/actors/deactivate", middleware.SuperadminMiddleware(), arh.Deactivate)

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
