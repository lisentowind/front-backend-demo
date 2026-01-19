package main

import (
	"backend/internal/router"
	"fmt"
)

func main() {
	r := router.InitRouter()

	port := ":8080"
	fmt.Printf("服务正在启动，监听端口 %s ...\n", port)

	if err := r.Run(port); err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}
