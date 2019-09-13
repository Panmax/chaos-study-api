package courses

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
)

type CourseModelValidator struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Total uint16 `form:"total" json:"total" binding:"required"`
	Url   string `form:"url" json:"url" binding:"max=2048"`
	Pick  uint8  `form:"pick" json:"pick"`

	courseModel CourseModel `json:"-"`
}

func NewCourseModelValidator() CourseModelValidator {
	return CourseModelValidator{}
}

func (v *CourseModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}

	v.courseModel.Name = v.Name
	v.courseModel.Total = v.Total
	v.courseModel.Pick = v.Pick
	if v.Url != "" {
		v.courseModel.Url = &v.Url
	}

	v.courseModel.UserId = 1 // FIXME
	return nil
}
