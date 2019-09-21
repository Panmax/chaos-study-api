package users

import (
	"errors"
	"github.com/Panmax/chaos-study-api/common"
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

	if err := CreateUser(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessResponse(true))
}

type passwordForm struct {
	Password string `form:"password" json:"password" binding:"exists,min=6,max=255"`
}

func UpdateUserPassword(c *gin.Context) {
	var form passwordForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	user := c.MustGet(common.JWTIdentityKey).(*UserModel)
	if err := user.setPassword(form.Password); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	if err := SaveOne(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(true))
}
