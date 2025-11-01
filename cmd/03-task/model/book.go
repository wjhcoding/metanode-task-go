package model

// Book 对应 books 表
type Book struct {
	ID     int     `db:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title  string  `db:"title" gorm:"column:title"`
	Author string  `db:"author" gorm:"column:author"`
	Price  float64 `db:"price" gorm:"column:price"`
}

// TableName 指定表名
func (Book) TableName() string {
	return "books"
}
