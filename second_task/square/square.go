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

func (sqr Square) End() Point {
	return Point{sqr.start.x + int(sqr.a), sqr.start.y - int(sqr.a)}
}

func (sqr Square) Perimeter() uint {
	return sqr.a * 4
}

func (sqr Square) Area() uint {
	return uint(math.Pow(float64(sqr.a), float64(2)))
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
