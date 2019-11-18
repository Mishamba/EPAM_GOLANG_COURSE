package main

import "fmt"

func Factorial(givenNumber uint) (result uint) {
	result = 1
	for i := givenNumber; i > 1; i-- {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(0))
}
