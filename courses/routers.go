package courses

import (
	"errors"
	"fmt"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func CoursesRegister(router *gin.RouterGroup) {
	router.POST("/course", CreateCourse)
	router.PUT("/course/:id", UpdateCourse)
	router.DELETE("/course/:id", DeleteCourse)
	router.GET("/courses", ListCourse)
	router.GET("/course/:id", GetCourse)
	router.GET("/courses/pick", PickCourse)
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

	c.JSON(http.StatusOK, common.NewSuccessResponse(nil))
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

	courseModelValidator := NewCourseModelValidator()
	if err := courseModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	courseModel.Name = courseModelValidator.Name
	courseModel.Chapters = courseModelValidator.Chapters
	courseModel.Url = courseModelValidator.Url
	courseModel.Pick = courseModelValidator.Pick
	if err := SaveOne(&courseModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(nil))
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

	c.JSON(http.StatusOK, common.NewSuccessResponse(nil))
}

func ListCourse(c *gin.Context) {
	limit := c.Query("limit")
	offset := c.Query("offset")

	courseModels, total, err := FindCourse(limit, offset)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	results := make([]CourseResponse, 0)
	for _, courseModel := range courseModels {
		serializer := CourseSerializer{courseModel}
		results = append(results, serializer.Response())
	}
	paginationResponse := common.Pagination{Total: total, Results: results}

	c.JSON(http.StatusOK, common.NewSuccessResponse(paginationResponse))
}

func GetCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(errors.New("无效id")))
		return
	}
	courseModel, err := FindOneCourse(uint(id)) // FIXME user filter
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	serializer := CourseSerializer{courseModel}
	c.JSON(http.StatusOK, common.NewSuccessResponse(serializer.Response()))
}

func PickCourse(c *gin.Context) {
	courseModels, err := FindAllCourse() // FIXME user filter
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	rand.Seed(time.Now().Unix())
	pickedCourse := courseModels[rand.Intn(len(courseModels))]
	section := rand.Intn(int(pickedCourse.Chapters))

	result := fmt.Sprintf("今日学习《%s》第 %d 节", pickedCourse.Name, section+1)
	c.JSON(http.StatusOK, common.NewSuccessResponse(result))
}
