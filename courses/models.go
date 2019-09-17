package courses

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/jinzhu/gorm"
	"strconv"
)

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

type CourseModel struct {
	gorm.Model

	UserId uint

	Name         string
	TotalChapter uint16
	Url          string
	Pick         uint8
}

func (CourseModel) TableName() string {
	return "course"
}

func DeleteCourseModel(condition interface{}) error {
	db := common.GetDB()
	return db.Where(condition).Delete(CourseModel{}).Error
}

func FindCourse(userId uint, limit, offset string) (courses []CourseModel, total uint32, err error) {
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
	db = db.Where("user_id = ?", userId)

	db.Model(&courses).Count(&total)
	err = db.Offset(offsetInt).Limit(limitInt).Find(&courses).Error

	return
}

func FindAllCourse(userId uint) (courses []CourseModel, err error) {
	db := common.GetDB()
	err = db.Where("user_id = ?", userId).Find(&courses).Error

	return
}

func FindOneCourse(id uint) (course CourseModel, err error) {
	db := common.GetDB()
	err = db.First(&course, id).Error

	return
}

type CourseFlowModel struct {
	gorm.Model

	UserId uint

	Results CoursePickResults `gorm:"type:json"`
}

func (CourseFlowModel) TableName() string {
	return "course_flow"
}

func FindTodayCourseFlow(userId uint) (flow CourseFlowModel, err error) {
	var courseFlow CourseFlowModel

	db := common.GetDB()
	err = db.Where("user_id = ?", userId).Where("created_at > ?", common.GetToday()).First(&courseFlow).Error
	return
}

func ExistCourseFlowByResult(userId uint, results CoursePickResults) (bool, error) {
	var count int
	db := common.GetDB()
	err := db.Model(&CourseFlowModel{}).Where("user_id = ?", userId).Where("results = ?", results).Count(&count).Error
	return count > 0, err
}
