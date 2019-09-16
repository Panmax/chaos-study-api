package users

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
)

type UserModelValidator struct {
	Username string `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
	Password string `form:"password" json:"password" binding:"exists,min=6,max=255"`
	Bio      string `form:"bio" json:"bio" binding:"max=1024"`

	userModel UserModel `json:"-"`
}

func (v *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	v.userModel.Username = v.Username
	v.userModel.Bio = v.Bio

	err = v.userModel.setPassword(v.Password)
	return err
}

func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	return userModelValidator
}
