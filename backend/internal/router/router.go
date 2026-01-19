package router

import (
	helloRouter "backend/internal/router/hello_router"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	helloRouter.InitHelloRouter(r)

	return r
}
