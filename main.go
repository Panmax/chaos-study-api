package main

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/courses"
	"github.com/Panmax/chaos-study-api/plans"
	"github.com/Panmax/chaos-study-api/users"
	"github.com/jinzhu/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&courses.CourseModel{})
	db.AutoMigrate(&courses.CourseFlowModel{})
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
		Bio:      "Go go go!",
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
	FillUpDB()
	defer db.Close()

	r := initRouter()
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}

}
