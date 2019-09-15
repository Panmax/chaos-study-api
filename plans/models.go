package plans

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/jinzhu/gorm"
)

type PlanModel struct {
	gorm.Model

	UserId uint `gorm:"unique_index"`

	Count     uint8
	NotRepeat bool
}

func (PlanModel) TableName() string {
	return "plan"
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindPlanByUser(userId uint) (PlanModel, error) {
	db := common.GetDB()
	var plan PlanModel

	err := db.Where("user_id = ?", userId).First(&plan).Error
	return plan, err
}
