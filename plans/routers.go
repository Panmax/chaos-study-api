package plans

import (
	"github.com/Panmax/chaos-study-api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SettingsRegister(router *gin.RouterGroup) {
	router.GET("/plan", GetPlan)
	router.PUT("/plan", UpdatePlan)
}

func GetPlan(c *gin.Context) {
	var userId uint = 1
	plan, err := FindPlanByUser(userId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	serializer := PlanSerializer{plan}
	c.JSON(http.StatusOK, common.NewSuccessResponse(serializer.Response()))
}

func UpdatePlan(c *gin.Context) {
	var userId uint = 1

	plan, err := FindPlanByUser(userId)
	if err != nil {
		plan = PlanModel{}
		plan.UserId = userId
	}

	planModelValidator := NewPlanModelValidator()
	err = planModelValidator.Bind(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}

	plan.Count = planModelValidator.planModel.Count
	plan.NotRepeat = planModelValidator.planModel.NotRepeat
	if err = SaveOne(&plan); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError(err))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(true))
}
