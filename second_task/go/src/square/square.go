package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (my_square Square) End() (x_result int, y_result int) {
	return my_square.start.x + int(my_square.a), my_square.start.y - int(my_square.a)
}

func (my_square Square) Perimeter() uint {
	return my_square.a * 4
}

//func (my_square Square) Area() uint {
//	return pow(my_square.a, 2)
//}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
//	fmt.Println(s.Area())
}
