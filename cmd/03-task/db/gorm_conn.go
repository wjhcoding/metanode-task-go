package db

import (
	"fmt"
	"log"

	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func InitGorm() {
	dsn := "root:jun123456@tcp(127.0.0.1:3306)/company?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ GORM 连接失败: " + err.Error())
	}
	GormDB = db
	fmt.Println("✅ GORM 数据库连接成功")

	// 自动迁移（根据模型自动创建或更新表结构）
	err = GormDB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		log.Fatalln("❌ 自动建表失败:", err)
	}
	fmt.Println("✅ 模型对应表已创建成功")
}
