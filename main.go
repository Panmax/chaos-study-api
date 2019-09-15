package main

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/courses"
	"github.com/Panmax/chaos-study-api/plans"
	"github.com/Panmax/chaos-study-api/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&courses.CourseModel{})
	db.AutoMigrate(&users.UserModel{})
	db.AutoMigrate(&plans.PlanModel{})
}

func FillUpDB() {
	db := common.GetDB()

	var userA users.UserModel
	if db.First(&userA); userA.ID != 0 { // 说明表非空
		return
	}

	tx1 := db.Begin()
	userA = users.UserModel{
		Username: "Panmax",
		Email:    "test@g.cn",
		Bio:      "Love Go",
		Image:    nil,
	}
	tx1.Save(&userA)

	planA := plans.PlanModel{
		UserId:    userA.ID,
		Count:     1,
		NotRepeat: true,
	}
	tx1.Save(&planA)
	tx1.Commit()
}

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	courses.CoursesRegister(v1)
	plans.SettingsRegister(v1)

	testAuth := r.Group("/api/ping")

	testAuth.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	FillUpDB()

	r.Run()
}
