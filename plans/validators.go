package plans

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
)

type PlanModelValidator struct {
	Count     uint8 `form:"count" json:"count" binding:"required"`
	NotRepeat *bool `form:"not_repeat" json:"not_repeat" binding:"exists"`

	planModel PlanModel `json:"-"`
}

func NewPlanModelValidator() PlanModelValidator {
	return PlanModelValidator{}
}

func (v *PlanModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	v.planModel.Count = v.Count
	v.planModel.NotRepeat = *v.NotRepeat

	v.planModel.UserId = 1 // FIXME
	return nil
}
