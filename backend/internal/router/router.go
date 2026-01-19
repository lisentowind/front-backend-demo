package router

import (
	"backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	helloController := &controller.HelloController{}

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", helloController.Ping)
	}

	return r
}
