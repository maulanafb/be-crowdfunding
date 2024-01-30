package main

import (
	"be_crowdfunding/handler"
	"be_crowdfunding/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AutoMigrateTables(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Database connected successfully")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userService.SaveAvatar(6, "images/avatar.jpg")
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()
}
