package courses

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"strings"
)

type CourseModelValidator struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Chapters uint16 `form:"chapters" json:"chapters" binding:"required"`
	Url      string `form:"url" json:"url" binding:"max=2048"`
	Pick     uint8  `form:"pick" json:"pick"`

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

	v.courseModel.Name = strings.TrimSpace(v.Name)
	v.courseModel.Chapters = v.Chapters
	v.courseModel.Pick = v.Pick
	v.courseModel.Url = strings.TrimSpace(v.Url)

	v.courseModel.UserId = 1 // FIXME
	return nil
}
