package main

import (
	"fmt"

	"sort"

	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/sorting/construct"
	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/sorting/tests"
	"github.com/Mishamba/EPAM_GOLANG_COURSE/FourthTask/sorting/types"
)

func main() {
	if err := tests.Errors(); err != nil {
		panic(err)
	} else {
		fmt.Println("All tests passed succesculy. Program can run")
	}
	fmt.Println("standart output")
	var somePersons types.People = types.People{
		types.Person{"g", "g", types.StartDate.AddDate(1990, 1, 1)},
		types.Person{"a", "a", types.StartDate.AddDate(1998, 1, 1)},
		types.Person{"z", "z", types.StartDate.AddDate(2000, 1, 1)},
		types.Person{"b", "b", types.StartDate.AddDate(1998, 1, 1)},
	}
	sort.Sort(somePersons)
	fmt.Println(somePersons)

	fmt.Println("manual input")

	People := construct.People()
	sort.Sort(People)
	fmt.Println(People)
}
