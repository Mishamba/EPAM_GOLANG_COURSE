package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(reverse(mySlice))
	newSlice := reverse(mySlice)
	fmt.Println(newSlice)
}

func reverse(givenSlice []int) (res []int) {
	res = make([]int, len(givenSlice))
	j := 0
	for i := len(givenSlice) - 1; i > -1; i-- {
		res[j] = givenSlice[i]
		j++
	}
	return res
}
