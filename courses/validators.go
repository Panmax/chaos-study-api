package courses

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"strings"
)

type CourseModelValidator struct {
	Name         string `form:"name" json:"name" binding:"required"`
	TotalChapter uint16 `form:"total_chapter" json:"total_chapter" binding:"required"`
	Url          string `form:"url" json:"url" binding:"max=2048"`
	Pick         uint8  `form:"pick" json:"pick"`

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
	v.courseModel.TotalChapter = v.TotalChapter
	v.courseModel.Pick = v.Pick
	v.courseModel.Url = strings.TrimSpace(v.Url)

	v.courseModel.UserId = c.MustGet(common.UserIDKey).(uint)
	return nil
}
