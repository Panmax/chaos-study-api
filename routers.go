package main

import (
	"fmt"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/courses"
	"github.com/Panmax/chaos-study-api/plans"
	"github.com/Panmax/chaos-study-api/users"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func initRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), handleErrors())

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
	return r
}
