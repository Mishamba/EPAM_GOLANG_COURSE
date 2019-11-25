package main

import (
	"fmt"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/figures/tests"
	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/figures/types"
)

func main() {
	if err := tests.Errors(); err != nil { //Programm will start work after tests check
		fmt.Print("Everthink is ok. Program can run.\n\n")
	}

	c, err := types.CircleConstruct(5) //Create circle with side 5.
	if err != nil {                    //Checking errors.
		panic(err) //If we have errors, we panic. we can't go ahead.
	}

	tmp, err := c.Area()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(tmp)
	}

	tmp, err = c.Perimetr()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(tmp)
	}

	sq, err := types.SquareConstruct(3) //Create square with radius 3.
	if err != nil {                     //Checking errors.
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
}
