package courses

import (
	"errors"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CoursesRegister(router *gin.RouterGroup) {
	router.POST("", CourseCreate)
	router.PUT("/:id", CourseUpdate)
	router.DELETE("/:id", CourseDelete)
	router.GET("", CourseList)
}

func CourseCreate(c *gin.Context) {
	courseModelValidator := NewCourseModelValidator()
	if err := courseModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	log.Println(courseModelValidator.courseModel)

	if err := SaveOne(&courseModelValidator.courseModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
	})
}

func CourseUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
	})
}

func CourseDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(errors.New("无效id")))
		return
	}
	err = DeleteCourseModel([]uint{uint(id)})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
	})
}

func CourseList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    []string{"1", "2", "3"},
	})
}
