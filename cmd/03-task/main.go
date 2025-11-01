package main

import (
	"fmt"
	"log"

	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/db"
	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/model"
	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/service"

	_ "github.com/go-sql-driver/mysql" // 使用MySQL驱动（换数据库需改此行）
)

func main() {
	// 初始化 SQLX 数据库
	db.InitSqlx()
	defer db.SqlxDB.Close()

	// ---------------------------
	// 1️⃣ 查询所有“技术部”员工
	// ---------------------------
	techEmployees, err := db.QueryTechEmployees()
	if err != nil {
		log.Fatal("查询技术部员工失败:", err)
	}
	fmt.Println("技术部员工列表：")
	for _, e := range techEmployees {
		fmt.Printf("ID:%d, 姓名:%s, 部门:%s, 工资:%.2f\n", e.ID, e.Name, e.Department, e.Salary)
	}

	// ---------------------------
	// 2️⃣ 查询工资最高的员工
	// ---------------------------
	topEmployee, err := db.QueryTopEmployee()
	if err != nil {
		log.Fatalln("查询工资最高的员工失败:", err)
	}

	fmt.Println("\n工资最高的员工：")
	fmt.Printf("ID:%d, 姓名:%s, 部门:%s, 工资:%.2f\n",
		topEmployee.ID, topEmployee.Name, topEmployee.Department, topEmployee.Salary)

	// ---------------------------
	// 3️⃣ 查询价格大于50元的书籍（类型安全）
	// ---------------------------
	expensiveBooks, err := db.QueryExpensiveBooks(50)
	if err != nil {
		log.Fatal("查询高价书籍失败:", err)
	}
	fmt.Println("\n价格大于50元的书籍列表：")
	for _, b := range expensiveBooks {
		fmt.Printf("ID:%d, 书名:%s, 作者:%s, 价格:%.2f\n", b.ID, b.Title, b.Author, b.Price)
	}

	// 初始化 GORM 数据库并自动建表
	db.InitGorm()

	// 插入示例数据
	service.SeedSampleData()

	// 查询指定用户的所有文章及评论
	service.QueryUserPostsWithComments(1)

	// 查询评论数最多的文章
	service.QueryMostCommentedPost()

	// 创建用户
	var user model.User
	if err := db.GormDB.Where("email = ?", "zhangsan@example.com").First(&user).Error; err != nil {
		user = model.User{Name: "张三", Email: "zhangsan@example.com"}
		db.GormDB.Create(&user)
	}

	// 创建文章（触发 AfterCreate）
	post := model.Post{Title: "GORM Hooks", Content: "学习GORM钩子函数", UserID: user.ID}
	db.GormDB.Create(&post)

	// 删除评论（触发 AfterDelete）
	comment1 := model.Comment{Content: "写得不错", User: "李四", PostID: post.ID}
	comment2 := model.Comment{Content: "有帮助", User: "王五", PostID: post.ID}
	db.GormDB.Create(&comment1)
	db.GormDB.Create(&comment2)

	// 删除两条评论（触发 AfterDelete）
	db.GormDB.Delete(&model.Comment{}, comment1.ID)
	db.GormDB.Delete(&model.Comment{}, comment2.ID)

	// 查看结果
	var updatedUser model.User
	var updatedPost model.Post
	db.GormDB.First(&updatedUser, user.ID)
	db.GormDB.First(&updatedPost, post.ID)

	fmt.Printf("\n👤 用户【%s】当前文章数：%d\n", updatedUser.Name, updatedUser.PostCount)
	fmt.Printf("📖 文章【%s】评论状态：%s\n", updatedPost.Title, updatedPost.CommentStatus)
}
