package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-backend-demo/internal/router/auth_router"
	helloRouter "go-backend-demo/internal/router/hello_router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有源，生产环境建议指定具体域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化认证路由
	auth_router.InitAuthRouter(r)

	// 初始化 Hello 路由
	helloRouter.InitHelloRouter(r)

	return r
}
