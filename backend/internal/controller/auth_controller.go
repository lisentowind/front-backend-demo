package controller

import (
	"net/http"

	"go-backend-demo/config"
	"go-backend-demo/internal/model"
	"go-backend-demo/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthController 认证控制器
type AuthController struct{}

// Login 用户登录
func (a *AuthController) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 查询用户
	var user model.User
	if err := config.DB.Where("name = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: http.StatusUnauthorized,
			Msg:  "用户名或密码错误",
			Data: nil,
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code: http.StatusUnauthorized,
			Msg:  "用户名或密码错误",
			Data: nil,
		})
		return
	}

	// 生成 Token
	token, err := utils.GenerateToken(user.Id, user.Name, "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "生成 Token 失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  "登录成功",
		Data: model.LoginResponse{
			Token: token,
		},
	})
}

// Register 用户注册
func (a *AuthController) Register(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	// 检查用户名是否已存在
	var existingUser model.User
	if err := config.DB.Where("name = ?", req.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: http.StatusBadRequest,
			Msg:  "用户名已存在",
			Data: nil,
		})
		return
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "密码加密失败",
			Data: nil,
		})
		return
	}

	// 创建用户
	user := model.User{
		Name:       req.Username,
		Password:   string(hashedPassword),
		CreateTime: utils.GetChinaTime(),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "注册失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  "注册成功",
		Data: gin.H{
			"id":       user.Id,
			"username": user.Name,
		},
	})
}

// GetUserInfo 获取当前用户信息
func (a *AuthController) GetUserInfo(c *gin.Context) {
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")
	role, _ := c.Get("role")

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  "获取用户信息成功",
		Data: gin.H{
			"id":       userID,
			"username": username,
			"role":     role,
		},
	})
}
