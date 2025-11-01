package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/wjhcoding/MetaNode/task-go/cmd/03-task/model"
)

var SqlxDB *sqlx.DB

// InitSqlx 初始化 sqlx 连接
func InitSqlx() {
	var err error
	dsn := "root:jun123456@tcp(127.0.0.1:3306)/company?charset=utf8mb4&parseTime=True&loc=Local"
	SqlxDB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("数据库连接失败:", err)
	}
	fmt.Println("✅ Sqlx 数据库连接成功")
}

// QueryTechEmployees 查询技术部员工
func QueryTechEmployees() ([]model.Employee, error) {
	var employees []model.Employee
	query := `SELECT id, name, department, salary FROM employees WHERE department = ?`
	err := SqlxDB.Select(&employees, query, "技术部")
	return employees, err
}

// QueryTopEmployee 查询工资最高的员工
func QueryTopEmployee() (model.Employee, error) {
	var emp model.Employee
	query := `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`
	err := SqlxDB.Get(&emp, query)
	return emp, err
}

// QueryExpensiveBooks 查询价格大于指定值的书籍
func QueryExpensiveBooks(minPrice float64) ([]model.Book, error) {
	var books []model.Book
	query := `SELECT id, title, author, price FROM books WHERE price > ?`
	err := SqlxDB.Select(&books, query, minPrice)
	return books, err
}
