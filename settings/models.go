package settings

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/jinzhu/gorm"
)

type SettingModel struct {
	gorm.Model

	UserId uint `gorm:"unique_index"`

	Count     uint8
	NotRepeat bool
}

func (SettingModel) TableName() string {
	return "setting"
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func FindSetting(userId uint) (SettingModel, error) {
	db := common.GetDB()
	var setting SettingModel

	err := db.Where("user_id = ?", userId).First(&setting).Error
	return setting, err
}
