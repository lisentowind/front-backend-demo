package model

type User struct {
	Id         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Age        int    `gorm:"type:int" json:"age"`
	Email      string `gorm:"type:varchar(100)" json:"email"`
	Password   string `gorm:"type:varchar(255);not null" json:"-"` // 密码字段，不返回给前端
	CreateTime string `gorm:"type:varchar(50)" json:"createTime"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应结构
type LoginResponse struct {
	Token string `json:"token"`
}

// 明确指定表名
func (User) TableName() string {
	return "user"
}
