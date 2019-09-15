package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
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
	if reflect.TypeOf(err) == reflect.TypeOf(validator.ValidationErrors{}) {
		return NewValidatorError(err)
	}

	res := CommonError{Code: -1}
	res.Message = err.Error()
	return res
}

func Bind(c *gin.Context, obj interface{}) error {
	return c.ShouldBind(obj)
}

func InSliceInt(sl []int, v int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func GetToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}
