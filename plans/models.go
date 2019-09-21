package plans

import (
	"github.com/Panmax/chaos-study-api/common"
)

type PlanModel struct {
	common.Model

	UserId uint `gorm:"unique_index;not null"`

	Count     uint8 `gorm:"not null"`
	NotRepeat bool  `gorm:"not null"`
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
