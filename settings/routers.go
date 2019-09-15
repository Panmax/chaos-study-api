package settings

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SettingsRegister(router *gin.RouterGroup) {
	router.GET("/settings", GetSetting)
	router.PUT("/settings", UpdateSetting)
}

func GetSetting(c *gin.Context) {
	var userId uint = 1
	setting, err := FindSetting(userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	serializer := SettingSerializer{setting}
	c.JSON(http.StatusOK, common.NewSuccessResponse(serializer.Response()))
}

func UpdateSetting(c *gin.Context) {
	var userId uint = 1

	setting, err := FindSetting(userId)
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
