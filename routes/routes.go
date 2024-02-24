package routes

import (
	"github.com/504BassSlapper/gin-experimentation/middleware"
	"github.com/gin-gonic/gin"
)

func GetMenuGroupRoutes(router *gin.Engine) *gin.Engine {
	routerGroup := router.Group("v1")
	{
		routerGroup.GET("/healthz", middleware.HealthCheck)
	}
	return router
}
