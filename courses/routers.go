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
	router.POST("", CreateCourse)
	router.PUT("/:id", UpdateCourse)
	router.DELETE("/:id", DeleteCourse)
	router.GET("", ListCourse)
	router.GET("/:id", GetCourse)
}

func CreateCourse(c *gin.Context) {
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

func UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(errors.New("无效id")))
		return
	}
	courseModel, err := FindOneCourse(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	courseModelValidator := NewCourseModelValidatorFillWith(courseModel)
	if err := courseModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	courseModelValidator.courseModel.ID = courseModel.ID
	if err := SaveOne(&courseModelValidator.courseModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
	})
}

func DeleteCourse(c *gin.Context) {
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

func ListCourse(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	courseModels, total, err := FindCourse(limit, offset)

	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	pageResponse := make(map[string]interface{})
	pageResponse["total"] = total
	pageResponse["results"] = courseModels
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    pageResponse,
	})
}

func GetCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(errors.New("无效id")))
		return
	}
	articleModel, err := FindOneCourse(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "OK",
		"data":    articleModel,
	})
}
