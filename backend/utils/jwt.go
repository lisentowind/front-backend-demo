package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 密钥 - 生产环境应该从环境变量读取
var jwtSecret = []byte("your-very-secure-secret-key-change-this-in-production")

// MyClaims 自定义声明结构体
type MyClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"` // "admin" 或 "user"
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID int, username string, role string) (string, error) {
	claims := MyClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间：2小时
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			// 签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
			// 签发人
			Issuer: "go-backend",
		},
	}

	// 使用 HS256 算法签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用 secret 签名
	return token.SignedString(jwtSecret)
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 校验签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
