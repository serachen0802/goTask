package router

import (
	"goTask/service"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", service.GetTask)
	router.POST("/task", service.Task)
	router.PUT("/task/:id", service.UpdateTask)
	router.DELETE("/task/:id", service.Delete)

	return router
}
