package users

import (
	"errors"
	"github.com/Panmax/chaos-study-api/common"
	"github.com/Panmax/chaos-study-api/plans"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersAnonymousRegister(router *gin.RouterGroup) {
	router.POST("/user", RegistrationUsers)
}

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/user/password", UpdateUserPassword)
}

func RegistrationUsers(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	_, err := FindOneUser(userModelValidator.userModel.Username)
	if err == nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(errors.New("用户名已存在")))
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}
	if err = plans.SaveOne(&plans.PlanModel{UserId: userModelValidator.userModel.ID, Count: 1, NotRepeat: true}); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessResponse(true))
}

func UpdateUserPassword(c *gin.Context) {
	c.String(http.StatusOK, "update password")
}
