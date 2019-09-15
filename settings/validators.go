package settings

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
)

type SettingModelValidator struct {
	Count     uint8 `form:"count" json:"count" binding:"required"`
	NotRepeat *bool `form:"not_repeat" json:"not_repeat" binding:"exists"`

	settingModel SettingModel `json:"-"`
}

func NewSettingModelValidator() SettingModelValidator {
	return SettingModelValidator{}
}

func (v *SettingModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	v.settingModel.Count = v.Count
	v.settingModel.NotRepeat = *v.NotRepeat

	v.settingModel.UserId = 1 // FIXME
	return nil
}
