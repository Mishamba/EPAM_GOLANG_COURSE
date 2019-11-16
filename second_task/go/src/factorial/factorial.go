package main

import "fmt"

func factorial(given_number uint) (result uint) { //могу ли я как-то инициализировать result прямо в шапке функции?
	result=1
	if given_number == 0 {
		return result
	}
	for i:=given_number ; i>1 ; i-- {
		result *= i
	}
	return result
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(0))
}
