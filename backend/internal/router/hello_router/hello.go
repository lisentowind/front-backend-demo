package helloRouter

import (
	"go-backend-demo/internal/controller"
	"go-backend-demo/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitHelloRouter(defaultRouter *gin.Engine) {
	helloController := &controller.HelloController{}

	// 公开路由组 - 不需要认证
	publicRouter := defaultRouter.Group("/api/v1/hello")
	{
		publicRouter.GET("/ping", helloController.Ping)
		publicRouter.GET("/user/table", helloController.HelloTable)
	}

	// 需要认证的路由组
	authRouter := defaultRouter.Group("/api/v1/hello")
	authRouter.Use(middleware.AuthMiddleware())
	{
		authRouter.GET("/user/all", helloController.GetAllUsers)
		authRouter.GET("/user", helloController.GetUserById)
		authRouter.POST("/user/add", helloController.AddUser)
		authRouter.DELETE("/user/delete", helloController.DeleteUsers)
	}
}
