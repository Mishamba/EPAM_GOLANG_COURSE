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
	if tmp, err = c.Area(); err != nil {
		panic(err)
	} else {
		fmt.Println(tmp)
	}

	if tmp, err = c.Area(); err != nil {
		panic(err)
	} else {
		fmt.Println(tmp)
	}

	sq, err := types.NewSquare(3) //Create square with radius 3.
	if err != nil {               //Checking errors.
		panic(err) //If we have errors, we panic. we can't go ahead.
	}

	tmp, err = sq.Area()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(tmp)
	}

	tmp, err = sq.Perimetr()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(tmp)
	}

	a, err := types.NewCircle(-3)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}
