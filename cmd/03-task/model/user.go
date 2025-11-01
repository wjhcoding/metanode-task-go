package model

import "gorm.io/gorm"

// User 用户模型：一个用户可以发表多篇文章（HasMany）
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null;comment:用户名"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null;comment:邮箱"`
	Password  string `gorm:"type:varchar(255);not null;comment:密码"`
	Posts     []Post `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // 一对多
	PostCount int64  `gorm:"default:0"`                                                        // 🆕 文章数量统计字段
}

func (User) TableName() string {
	return "users"
}
