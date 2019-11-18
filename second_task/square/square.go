package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (MySquare Square) End() (XResult int, YResult int) {
	return MySquare.start.x + int(MySquare.a), MySquare.start.y - int(MySquare.a)
}

func (MySquare Square) Perimeter() uint {
	return MySquare.a * 4
}

func (MySquare Square) Area() uint {
	return uint(math.Pow(float64(MySquare.a), float64(2)))
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
