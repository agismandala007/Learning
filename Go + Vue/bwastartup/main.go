package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// fmt.Println("connection success!!")

	// var users []user.User
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// }

	// router := gin.Default()
	// router.GET("/", handler)

	// router.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println("connection success!!")

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)
// }
