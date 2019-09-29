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

const MaxFindCount int = 20

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
	if courseModel.UserId != c.MustGet(common.UserIDKey).(uint) {
		c.JSON(http.StatusForbidden, common.NewError(errors.New("无权限")))
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
	courseModel, err := FindOneCourse(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}
	if courseModel.UserId != c.MustGet(common.UserIDKey).(uint) {
		c.JSON(http.StatusForbidden, common.NewError(errors.New("无权限")))
	}

	DeleteCourseModel([]uint{uint(id)})
	c.JSON(http.StatusOK, common.NewSuccessResponse(true))
}

func ListCourse(c *gin.Context) {
	userId := c.MustGet(common.UserIDKey).(uint)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))

	courseModels, total, err := FindCourse(userId, page, size)
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
	courseModel, err := FindOneCourse(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError(err))
		return
	}

	serializer := CourseSerializer{courseModel}
	c.JSON(http.StatusOK, common.NewSuccessResponse(serializer.Response()))
}

func PickCourse(c *gin.Context) {
	userId := c.MustGet(common.UserIDKey).(uint)

	flows, err := FindCourseFlow(userId, "created_at > ?", common.GetToday())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	} else if len(flows) == 0 {
		plan, err := plans.FindPlanByUser(userId)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
			return
		}

		flows, _ = pickCourse(userId, plan)
		for _, flow := range flows {
			SaveOne(&flow)
		}
	}

	c.JSON(http.StatusOK, makePickResponse(flows))
}

func pickCourse(userId uint, plan plans.PlanModel) ([]CourseFlowModel, error) {

	courseModels, err := FindAllCourse(userId)
	if err != nil {
		return nil, err
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

	var results []CourseFlowModel
	for _, pickedCourse := range pickedCourses {
		var chapters []int
		maxTry := int(pickedCourse.Pick * 5)
		tryCounter := 0
		for {
			tryCounter++
			if tryCounter > maxTry {
				break
			}

			chapter := rand.Intn(int(pickedCourse.TotalChapter)) + 1 // 从第一节开始

			if plan.NotRepeat {
				if exist, err := ExistCourseFlowByCourseAndChapter(userId, pickedCourse.ID, uint16(chapter)); err != nil {
					return nil, err
				} else if exist {
					continue
				}
			}

			chapters = append(chapters, chapter)
			if len(chapters) >= int(pickedCourse.Pick) || len(chapters) >= int(pickedCourse.TotalChapter) {
				break
			}
		}
		sort.Ints(chapters)

		for _, chapter := range chapters {
			results = append(results, CourseFlowModel{UserId: userId, CourseId: pickedCourse.ID, Chapter: uint16(chapter)})
		}

	}

	return results, nil
}

func makePickResponse(flows []CourseFlowModel) []*CoursePickResponse {
	var results []*CoursePickResponse
	for _, flow := range flows {

		find := false
		for _, r := range results {
			if r.Course.ID == flow.CourseId {
				r.Chapters = append(r.Chapters, int(flow.Chapter))
				find = true
				break
			}
		}
		if find {
			continue
		}

		var response CoursePickResponse
		course, _ := FindOneCourse(flow.CourseId)
		courseSerializer := CourseSerializer{course}
		response.Course = courseSerializer.Response()
		response.Chapters = append(response.Chapters, int(flow.Chapter))
		results = append(results, &response)
	}
	return results
}
