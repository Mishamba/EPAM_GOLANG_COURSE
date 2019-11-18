package main

import "fmt"

func Factorial(GivenNumber uint) (Result uint) {
	Result = 1
	if GivenNumber == 0 {
		return Result
	}
	for i := GivenNumber; i > 1; i-- {
		Result *= i
	}
	return Result
}

func main() {
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(0))
}
