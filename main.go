package main

import (
	"fmt"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/courses"
	"github.com/Panmax/chaos-study-api/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&courses.CourseModel{})
	db.AutoMigrate(&users.UserModel{})
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	courses.CoursesRegister(v1.Group("/courses"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	tx1 := db.Begin()
	userA := users.UserModel{
		Username: "Panmax",
		Email:    "test@g.cn",
		Bio:      "Love Go",
		Image:    nil,
	}
	tx1.Save(&userA)
	tx1.Commit()
	fmt.Println(userA)

	r.Run()
}
