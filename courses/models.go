package courses

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/jinzhu/gorm"
	"strconv"
)

type CourseModel struct {
	gorm.Model

	UserId uint

	Name  string
	Total uint16
	Url   string
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

func FindCourse(limit, offset string) ([]CourseModel, uint32, error) {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
		err = nil
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 20
		err = nil
	}

	db := common.GetDB()
	var courses []CourseModel
	var total uint32

	db.Model(&courses).Count(&total)
	err = db.Offset(offsetInt).Limit(limitInt).Find(&courses).Error

	return courses, total, err
}

func FindAllCourse() ([]CourseModel, error) {
	var courses []CourseModel

	db := common.GetDB()
	err := db.Find(&courses).Error

	return courses, err
}

func FindOneCourse(id uint) (CourseModel, error) {
	db := common.GetDB()
	var course CourseModel
	err := db.First(&course, id).Error
	return course, err
}
