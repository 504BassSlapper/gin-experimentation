package main

import (
	"fmt"
	"net/http"

	"github.com/504BassSlapper/gin-experimentation/middleware"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router

}

func getV1Router() *gin.Engine {
	router := gin.Default()
	routerGroup := router.Group("/v1")
	routerGroup.GET("/healthz", middleware.HealthCheck)
	return router

}

func main() {
	r := getV1Router()
	fmt.Println(r)
	r.Run(":5000")
}
