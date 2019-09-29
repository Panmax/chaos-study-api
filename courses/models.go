package courses

import (
	"github.com/Panmax/chaos-study-api/common"
	"strconv"
	"time"
)

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

type CourseModel struct {
	common.Model

	UserId uint `gorm:"not null"`

	Name         string `gorm:"size:128;not null"`
	TotalChapter uint16 `gorm:"not null"`
	Url          string `gorm:"not null"`
	Pick         uint8  `gorm:"not null"`
	DeletedAt    *time.Time
}

func (CourseModel) TableName() string {
	return "course"
}

type CourseFlowModel struct {
	common.Model

	UserId uint `gorm:"not null"`

	CourseId uint   `gorm:"not null"`
	Chapter  uint16 `gorm:"not null"`
}

func (CourseFlowModel) TableName() string {
	return "course_flow"
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
	err = db.Unscoped().First(&course, id).Error

	return
}

func FindCourseFlow(userId uint, args ...interface{}) (flows []CourseFlowModel, err error) {
	db := common.GetDB().Where("user_id = ?", userId)
	if len(args) >= 2 {
		db = db.Where(args[0], args[1:]...)
	} else if len(args) >= 1 {
		db = db.Where(args[0])
	}
	err = db.Find(&flows).Error

	return
}

func ExistCourseFlowByCourseAndChapter(userId uint, courseId uint, chapter uint16) (bool, error) {
	var count int
	db := common.GetDB()
	err := db.Model(&CourseFlowModel{}).Where("user_id = ?", userId).Where(
		"course_id = ?", courseId).Where("chapter = ?", chapter).Count(&count).Error
	return count > 0, err
}
