package main

import (
	"fmt"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/courses"
	"github.com/Panmax/chaos-study-api/plans"
	"github.com/Panmax/chaos-study-api/users"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
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
	defer db.Close()

	r := gin.Default()

	authMiddleware, err := users.NewGinJWTMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	v1 := r.Group("/api")
	v1.POST("/auth/login", authMiddleware.LoginHandler)
	v1.POST("/auth/refresh_token", authMiddleware.RefreshHandler)
	users.UsersAnonymousRegister(v1)

	v1.Use(authMiddleware.MiddlewareFunc())
	users.UsersRegister(v1)
	courses.CoursesRegister(v1)
	plans.PlansRegister(v1)

	testAuth := r.Group("/api/ping")
	testAuth.Use(authMiddleware.MiddlewareFunc())

	testAuth.GET("", func(c *gin.Context) {
		user := c.MustGet(common.JWTIdentityKey).(*users.UserModel)
		fmt.Println(user)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong, " + user.Username,
		})
	})

	r.GET("", func(c *gin.Context) {
		c.String(http.StatusOK, "Chaos Study API")
	})

	FillUpDB()

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}

}
