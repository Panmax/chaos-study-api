package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

type CommonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewValidatorError(err error) CommonError {
	res := CommonError{Code: -1, Message: "参数错误"}
	res.Data = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		if v.Param != "" {
			res.Data.(map[string]interface{})[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Data.(map[string]interface{})[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

func NewError(err error) CommonError {
	res := CommonError{Code: -1}
	res.Message = err.Error()
	return res
}

func Bind(c *gin.Context, obj interface{}) error {
	return c.ShouldBind(obj)
}
