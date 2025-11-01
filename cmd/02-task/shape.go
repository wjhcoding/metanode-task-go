package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func RunShapeTask() {
	// 新增：在主函数中创建 Rectangle 和 Circle，并调用它们的 Area() 和 Perimeter()
	var s Shape

	rect := Rectangle{Width: 5, Height: 10}
	s = rect
	fmt.Printf("矩形面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())

	circ := Circle{Radius: 7}
	s = circ
	fmt.Printf("圆面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}
