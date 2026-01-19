package model

type User struct {
	Id         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Age        int    `gorm:"type:int" json:"age"`
	Email      string `gorm:"type:varchar(100)" json:"email"`
	CreateTime string `gorm:"type:varchar(50)" json:"createTime"`
}

// 明确指定表名
func (User) TableName() string {
	return "user"
}
