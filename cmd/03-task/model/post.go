package model

import (
	"fmt"

	"gorm.io/gorm"
)

// Post 文章模型：一篇文章属于一个用户（BelongsTo）
// 一篇文章有多个评论（HasMany）
type Post struct {
	gorm.Model
	Title    string    `gorm:"type:varchar(100);not null;comment:标题"`
	Content  string    `gorm:"type:text;not null;comment:内容"`
	UserID   uint      `gorm:"index;comment:作者ID"` // 外键
	Comments []Comment `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CommentStatus string `gorm:"size:20;default:'有评论'"` // 🆕 评论状态
}

func (Post) TableName() string {
	return "posts"
}

/**
 * AfterCreate 在 Create() 操作成功后自动调用；
 * 使用 gorm.Expr 确保并发更新安全；
 * tx 是当前事务上下文。
 */
// AfterCreate 钩子：在文章创建后执行
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新对应用户的文章数量
	err = tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + ?", 1)).
		Error

	if err != nil {
		fmt.Printf("❌ 更新用户文章数量失败: %v\n", err)
		return err
	}

	fmt.Printf("✅ 用户ID=%d 的文章数量已增加\n", p.UserID)
	return nil
}
