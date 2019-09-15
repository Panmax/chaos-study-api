package courses

import (
	"errors"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/plans"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"sort"
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

	c.JSON(http.StatusOK, common.NewSuccessResponse(true))
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
	courseModel.TotalChapter = courseModelValidator.TotalChapter
	courseModel.Url = courseModelValidator.Url
	courseModel.Pick = courseModelValidator.Pick
	if err := SaveOne(&courseModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(true))
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

	c.JSON(http.StatusOK, common.NewSuccessResponse(true))
}

func ListCourse(c *gin.Context) {
	var userId uint = 1 // FIXME

	limit := c.Query("limit")
	offset := c.Query("offset")

	courseModels, total, err := FindCourse(userId, limit, offset)
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
	var userId uint = 1 // FIXME

	courseFlow, err := FindTodayCourseFlow(userId)
	if err == nil {
		c.JSON(http.StatusOK, common.NewSuccessResponse(courseFlow.Results))
		return
	}

	plan, err := plans.FindPlanByUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	courseModels, err := FindAllCourse(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	count := int(plan.Count)
	if count > len(courseModels) {
		count = len(courseModels)
	}

	rand.Seed(time.Now().Unix())
	var pickedCourses []CourseModel
	var index int
	for i := 0; i < int(count); i++ {
		index = rand.Intn(len(courseModels))
		pickedCourses = append(pickedCourses, courseModels[index])
		courseModels = append(courseModels[:index], courseModels[index+1:]...) // 将挑选出的课程从列表中移除
	}

	var results []CoursePickResponse
	for _, pickedCourse := range pickedCourses {
		var chapters []int
		for {
			chapter := rand.Intn(int(pickedCourse.TotalChapter)) + 1 // 从第一节开始
			if !common.InSliceInt(chapters, chapter) {
				chapters = append(chapters, chapter)
			}
			if len(chapters) >= int(pickedCourse.Pick) || len(chapters) >= int(pickedCourse.TotalChapter) {
				break
			}
		}
		sort.Ints(chapters)

		courseSerializer := CourseSerializer{pickedCourse}
		results = append(results, CoursePickResponse{Course: courseSerializer.Response(), Chapters: chapters})
	}

	err = SaveOne(&CourseFlowModel{UserId: userId, Results: results})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(results))
}
