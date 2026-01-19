package helloRouter

import (
	"backend/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitHelloRouter(defaultRouter *gin.Engine) {
	helloController := &controller.HelloController{}
	helloRouter := defaultRouter.Group("/api/v1/hello")
	{
		helloRouter.GET("/ping", helloController.Ping)
		helloRouter.GET("/user/table", helloController.HelloTable)
		helloRouter.GET("/user/all", helloController.GetAllUsers)
		helloRouter.GET("/user", helloController.GetUserById)
		helloRouter.POST("/user/add", helloController.AddUser)
	}
}
