package middleware

import (
	"net/http"
	"strings"

	"go-backend-demo/internal/model"
	"go-backend-demo/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取请求头中的 Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Msg:  "请求头中缺少 Authorization",
			})
			c.Abort()
			return
		}

		// 2. 按格式提取 Token (格式: "Bearer {token}")
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Msg:  "Authorization 格式有误，应为: Bearer {token}",
			})
			c.Abort()
			return
		}

		// 3. 解析 Token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Msg:  "无效的 Token",
			})
			c.Abort()
			return
		}

		// 4. 将用户信息保存到请求上下文
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// AdminAuthMiddleware 管理员权限中间件
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先执行认证
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Msg:  "请求头中缺少 Authorization",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Msg:  "Authorization 格式有误",
			})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Response{
				Code: http.StatusUnauthorized,
				Data: nil,
				Msg:  "无效的 Token",
			})
			c.Abort()
			return
		}

		// 检查角色是否为管理员
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, model.Response{
				Code: http.StatusForbidden,
				Data: nil,
				Msg:  "权限不足，需要管理员权限",
			})
			c.Abort()
			return
		}

		// 保存用户信息
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
