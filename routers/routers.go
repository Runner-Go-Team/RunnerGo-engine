package routers

import (
	"github.com/gin-gonic/gin"
	"kp-runner/api"
)

func InitRouter(Router *gin.RouterGroup) {
	{
		Router.POST("/run_plan/", api.RunPlan)
		Router.POST("/run_api/", api.RunApi)
		Router.POST("/run_scene/", api.RunScene)
		Router.GET("/get_id", api.GetId)
	}
}
