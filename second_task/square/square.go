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

func (Sqr Square) End() Point {
	return Point{Sqr.start.x + int(Sqr.a), Sqr.start.y - int(Sqr.a)}
}

func (Sqr Square) Perimeter() uint {
	return Sqr.a * 4
}

func (Sqr Square) Area() uint {
	return uint(math.Pow(float64(Sqr.a), float64(2)))
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
