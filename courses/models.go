package courses

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/jinzhu/gorm"
)

type CourseModel struct {
	gorm.Model

	UserId uint

	Name  string
	Total uint16
	Url   *string
	Pick  uint8
}

func (CourseModel) TableName() string {
	return "course"
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func DeleteCourseModel(condition interface{}) error {
	db := common.GetDB()
	err := db.Where(condition).Delete(CourseModel{}).Error
	return err
}
