package auth_router

import (
	"go-backend-demo/internal/controller"
	"go-backend-demo/internal/middleware"

	"github.com/gin-gonic/gin"
)

// InitAuthRouter 初始化认证路由
func InitAuthRouter(r *gin.Engine) {
	authController := &controller.AuthController{}
	authGroup := r.Group("/api/v1/auth")
	{
		// 公开路由 - 不需要认证
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)

		// 需要认证的路由
		authGroup.Use(middleware.AuthMiddleware())
		{
			authGroup.GET("/info", authController.GetUserInfo)
		}
	}
}
