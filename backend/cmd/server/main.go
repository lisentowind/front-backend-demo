package main

import (
	"backend/config"
	"backend/internal/router"
	"fmt"
)

func main() {
	// 连接数据库
	config.InitMySQL()

	// 初始化路由
	r := router.InitRouter()

	port := ":8080"
	fmt.Printf("服务正在启动，监听端口 %s ...\n", port)

	if err := r.Run(port); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}
