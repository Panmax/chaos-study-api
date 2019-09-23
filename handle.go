package main

import (
	"errors"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(http.StatusInternalServerError,
					common.NewError(errors.New("服务器开小差了…")))
				return
			}
		}()
		c.Next()
	}
}
