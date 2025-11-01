package model

// Employee 对应 employees 表
type Employee struct {
	ID         int     `db:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name       string  `db:"name" gorm:"column:name"`
	Department string  `db:"department" gorm:"column:department"`
	Salary     float64 `db:"salary" gorm:"column:salary"`
}

// TableName 指定表名（用于 GORM）
func (Employee) TableName() string {
	return "employees"
}
