package service

import (
	"fmt"

	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/db"
	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/model"
	"gorm.io/gorm"
)

// QueryUserPostsWithComments 查询某个用户发布的所有文章及其评论
func QueryUserPostsWithComments(userID uint) {
	var user model.User

	err := db.GormDB.Preload("Posts.Comments"). // 预加载文章及每篇文章的评论
							First(&user, userID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("❌ 未找到ID=%d的用户\n", userID)
			return
		}
		panic(err)
	}

	fmt.Printf("✅ 用户: %s 的文章及评论：\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("\n📖 文章: %s\n内容: %s\n", post.Title, post.Content)
		if len(post.Comments) == 0 {
			fmt.Println("  （暂无评论）")
		} else {
			for _, c := range post.Comments {
				fmt.Printf("  💬 评论者: %s → %s\n", c.User, c.Content)
			}
		}
	}
}

// QueryMostCommentedPost 查询评论数量最多的文章
func QueryMostCommentedPost() {
	type Result struct {
		PostID       uint
		Title        string
		CommentCount int64
	}

	var result Result

	err := db.GormDB.
		Table("posts p").
		Select("p.id as post_id, p.title, COUNT(c.id) as comment_count").
		Joins("LEFT JOIN comments c ON p.id = c.post_id").
		Group("p.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error

	if err != nil {
		panic(err)
	}

	if result.PostID == 0 {
		fmt.Println("⚠️ 暂无文章或评论数据")
		return
	}

	fmt.Printf("\n🔥 评论最多的文章：\nID: %d\n标题: %s\n评论数量: %d\n",
		result.PostID, result.Title, result.CommentCount)
}
