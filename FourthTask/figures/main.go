package main

import (
	"fmt"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/figures/types"
)

func main() {
	c, err := types.NewCircle(5) //Create circle with side 5.
	if err != nil {              //Checking errors.
		panic(err) //If we have errors, we panic. we can't go ahead.
	}

	var tmp float64
	tmp = c.Area()
	fmt.Println(tmp)

	tmp = c.Area()
	fmt.Println(tmp)

	sq, err := types.NewSquare(3) //Create square with radius 3.
	if err != nil {               //Checking errors.
		panic(err) //If we have errors, we panic. we can't go ahead.
	}

	tmp = sq.Area()
	fmt.Println(tmp)

	tmp = sq.Perimetr()
	fmt.Println(tmp)

	if a, err := types.NewCircle(-3); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}
}
