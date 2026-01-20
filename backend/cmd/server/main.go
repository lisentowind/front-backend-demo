package main

import (
	"fmt"

	"go-backend-demo/config"
	"go-backend-demo/internal/router"
)

func main() {
	// 连接数据库
	config.InitMySQL()

	// 初始化路由
	r := router.InitRouter()

	addr := "0.0.0.0:8080"
	fmt.Printf("服务正在启动，监听地址 %s ...\n", addr)
	fmt.Println("可以通过以下方式访问:")
	fmt.Println("  - http://localhost:8080")
	fmt.Println("  - http://127.0.0.1:8080")
	fmt.Println("  - http://<本机IP>:8080")

	if err := r.Run(addr); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}
