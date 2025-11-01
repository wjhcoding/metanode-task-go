package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     // 匿名字段（组合）
	EmployeeID string
}

// PrintInfo 方法：输出员工信息
func (e Employee) PrintInfo() {
	fmt.Printf("员工信息：\n")
	fmt.Printf("姓名: %s\n", e.Name)
	fmt.Printf("年龄: %d\n", e.Age)
	fmt.Printf("员工ID: %s\n", e.EmployeeID)
}

func RunEmployeeTask() {
	// 输出员工的信息
	emp := Employee{
		Person: Person{
			Name: "张三",
			Age:  30,
		},
		EmployeeID: "E12345",
	}
	emp.PrintInfo()
}
