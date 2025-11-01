package service

import (
	"fmt"

	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/db"
	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/model"
)

// SeedSampleData 初始化测试数据
func SeedSampleData() {
	gormDB := db.GormDB

	// 检查是否已有数据，避免重复插入
	var count int64
	gormDB.Model(&model.User{}).Count(&count)
	if count > 0 {
		fmt.Println("⚠️ 示例数据已存在，跳过插入。")
		return
	}

	// 创建用户
	user := model.User{Name: "张三", Email: "zhangsan@example.com", Password: "123456"}
	gormDB.Create(&user)

	// 创建文章
	post := model.Post{Title: "GORM 一对多关系详解", Content: "本文介绍了 GORM 的一对多关系。", UserID: user.ID}
	gormDB.Create(&post)

	// 创建评论
	comments := []model.Comment{
		{Content: "写得很清晰！", PostID: post.ID, User: "李四"},
		{Content: "学到了，感谢！", PostID: post.ID, User: "王五"},
	}
	gormDB.Create(&comments)

	fmt.Println("✅ 示例数据插入成功！")
}
