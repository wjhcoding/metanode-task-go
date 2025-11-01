package model

import (
	"fmt"

	"gorm.io/gorm"
)

// Comment 评论模型：一条评论属于一篇文章
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null;comment:评论内容"`
	PostID  uint   `gorm:"index;not null;comment:所属文章ID"`
	User    string `gorm:"type:varchar(50);not null;comment:评论者"`
}

func (Comment) TableName() string {
	return "comments"
}

// AfterDelete 钩子：在评论删除后执行
// AfterDelete 在 Delete() 成功后触发；
// 在删除评论后统计同一文章的剩余评论数；
// 若为 0，则修改文章的 CommentStatus。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		fmt.Printf("❌ 统计评论失败: %v\n", err)
		return err
	}

	if count == 0 {
		err = tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").
			Error
		if err != nil {
			fmt.Printf("❌ 更新文章评论状态失败: %v\n", err)
			return err
		}

		fmt.Printf("⚠️ 文章ID=%d 的评论数为0，状态已更新为'无评论'\n", c.PostID)
	} else {
		fmt.Printf("📝 文章ID=%d 仍有 %d 条评论\n", c.PostID, count)
	}
	return nil
}
