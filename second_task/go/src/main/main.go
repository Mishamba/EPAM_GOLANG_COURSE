package main

import (
	"fmt"
	//	"math"
	"structures"
)

func main() {
	s := structures.Square{structures.Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
