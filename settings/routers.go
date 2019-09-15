package settings

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SettingsRegister(router *gin.RouterGroup) {
	router.PUT("/settings", UpdateSetting)
}

func UpdateSetting(c *gin.Context) {
	var userId uint = 1

	setting, err := FindSettingByUserId(userId)
	if err != nil {
		setting = SettingModel{}
		setting.UserId = userId
	}

	settingModelValidator := NewSettingModelValidator()
	err = settingModelValidator.Bind(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	setting.Count = settingModelValidator.settingModel.Count
	setting.NotRepeat = settingModelValidator.settingModel.NotRepeat
	if err = SaveOne(&setting); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(nil))
}